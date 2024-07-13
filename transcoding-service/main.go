package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"transcoding-service/rabbitmq"
	"transcoding-service/utils"
)

const (
	connUrl   string = "amqp://ajay:ajay@13.232.108.179:5672/"
	queueName string = "myqueue"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	go func() {
		<-sig
		done <- true
	}()

	mq := rabbitmq.NewRabbitmqClient(connUrl)

	messages := []rabbitmq.Message{
		{Data: "1"},
		{Data: "2"},
		{Data: "3"},
		{Data: "4"},
		{Data: "5"},
		{Data: "6"},
		{Data: "7"},
		{Data: "8"},
		{Data: "9"},
		{Data: "10"},
	}
	mq.OpenQueue(queueName)
	//
	consume, err := mq.Consume()
	utils.PanicIfError(err, "Unable to consume messages")
	//
	rabbitmq.PublishMessages(mq, messages)
	//
	time.Sleep(10 * time.Second)
	go rabbitmq.ConsumeMessages(consume)
	//
	//
	<-done
	time.Sleep(1 * time.Second)
	mq.Close()
	//
}
