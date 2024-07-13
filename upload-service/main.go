package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"upload-service/rabbitmq"
	"upload-service/routes"
	"upload-service/utils"

	"github.com/gin-gonic/gin"
)

const (
	HostUrl          = ":8081"
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

	//
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	group := r.Group("/api/v1")
	routes.RegisterRoutes(group)

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
