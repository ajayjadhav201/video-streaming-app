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
	connUrl   string = ""
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

	mq.QueueDeclare(queueName)
	//
	consume, err := mq.Consume()
	utils.PanicIfError(err, "Unable to consume messages")
	//
	//
	go func() {

		for msg := range consume {
			//
		}
	}()
	//
	//
	<-done
	time.Sleep(1 * time.Second)
	mq.Close()
	//
}
