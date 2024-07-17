package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"transcoding-service/rabbitmq"
	"transcoding-service/transcoder"
	"transcoding-service/utils"

	"github.com/rabbitmq/amqp091-go"
)

const (
	connUrl         string = ""
	queueName       string = "transcoding_service"
	VideoOutputPath string = "/path/to/output/"
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
	mq.QueueDeclare(queueName)
	//
	messages, err := mq.Consume()
	utils.PanicIfError(err, "Unable to consume messages")
	//
	//
	go conusmeMessages(mq, messages)
	//
	//
	<-done
	time.Sleep(1 * time.Second)
	mq.Close()
	//
}

func conusmeMessages(client *rabbitmq.RabbitmqClient, messages <-chan amqp091.Delivery) {
	//
	for msg := range messages {
		//
		var req rabbitmq.Message
		err := utils.Unmarshal(msg.Body, &req)
		if err != nil {
			utils.LogPrintln("Unable to decode message ", string(msg.Body), "error is: ", err.Error())
			continue
		}
		//process the video
		err = transcoder.TranscodeVideo(&req)
		if err != nil {
			utils.LogPrintln("Unable to decode message ", string(msg.Body), "error is: ", err.Error())
			continue
		}
		success := rabbitmq.Response{Message: "Success", OutputPath: VideoOutputPath + req.Title}
		data, err := utils.Marshal(&success)
		if err != nil {
			utils.LogPrintln("Unable to send success message ", string(msg.Body), "error is: ", err.Error())
			continue
		}
		client.PublishMessage("success", data)
	}
	//
}
