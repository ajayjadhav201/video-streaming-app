package controller

import (
	"upload-service/rabbitmq"
	"upload-service/utils"
)

var (
	MQ *rabbitmq.RabbitmqClient
)

func Init(client *rabbitmq.RabbitmqClient) {
	MQ = client
	utils.Printf("Controller initiated")
}
