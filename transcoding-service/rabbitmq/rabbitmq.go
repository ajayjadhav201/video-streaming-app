package rabbitmq

import (
	"bytes"
	"context"
	"time"
	"transcoding-service/utils"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitmqClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   *Queue
}

type Queue struct {
	queue amqp.Queue
}

func NewRabbitmqClient(connUrl string) *RabbitmqClient {
	conn, err := amqp.Dial(connUrl)
	utils.FatalIfError(err, "Failed to connect to rabbitmq: %s", connUrl)
	//
	// open a channel
	ch, err := conn.Channel()
	utils.PanicIfError(err, "Failed to open a channel")
	//
	utils.LogPrintf("Connected to rabbitmq")
	return &RabbitmqClient{
		conn:    conn,
		channel: ch,
		queue:   nil,
	}
}

func (c *RabbitmqClient) Close() {
	utils.LogPrintf("Closing rabbitmq connection...")
	c.conn.Close()
}

func (c *RabbitmqClient) OpenQueue(name string) *Queue {
	// declare a queue
	q, err := c.channel.QueueDeclare(
		name,
		true,  //durable
		false, //delete when unused
		false, //exclusive
		false, //no wait
		nil,   //arguments
	)
	utils.PanicIfError(err, "Failed to declare %s queue", name)
	// set Qos
	err = c.channel.Qos(
		1,     //prefetch count
		0,     // prefetch size
		false, //global
	)
	utils.PanicIfError(err, "Failed to declare set Qos")
	//
	queue := &Queue{
		queue: q,
	}
	c.queue = queue
	return queue
}

func (c *RabbitmqClient) PublishMessage(message []byte) error {
	// Publish message
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := c.channel.PublishWithContext(ctx,
		"",                 //exchange
		c.queue.queue.Name, //queueName
		false,              //mandetory
		false,              //immediate
		amqp.Publishing{
			// ContentType: "text/plain",
			Body: message,
		},
	)
	// returnChannel := make(chan amqp.Return)
	// c.channel.NotifyReturn(returnChannel)
	return err
}

func (c *RabbitmqClient) Consume() (<-chan amqp.Delivery, error) {
	messages, err := (c.channel).Consume(
		c.queue.queue.Name, //queueName
		"",                 // consumer name
		false,              // auto-ack
		false,              //exclusive
		false,              //no-local
		false,              //no-wait
		nil,                //args
	)
	// utils.PanicIfError(err, "Failed to register a consumer")
	return messages, err
}

func ConsumeMessages(messages <-chan amqp.Delivery) {
	for msg := range messages {
		utils.LogPrintf("Received a message: %s", msg.Body)
		dotCount := bytes.Count(msg.Body, []byte("."))
		t := time.Duration(dotCount)
		time.Sleep(t * time.Second)
		utils.LogPrintf("Done: %d", msg.DeliveryTag)
		msg.Ack(false)
	}
}

func PublishMessages(c *RabbitmqClient, messages []Message) {
	for _, msg := range messages {
		data, err := utils.Marshal(msg)
		if err != nil {
			utils.LogPrintf("Unable to marshal message: %s", err.Error())
			continue
		}
		utils.LogPrintf("Publishing a message: %s", msg)
		c.PublishMessage(data)
	}
}
