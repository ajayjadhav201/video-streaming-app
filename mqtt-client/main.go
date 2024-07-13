package main

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// MQTT broker address and client ID
	broker := "tcp://localhost:1883"
	clientID := "go_mqtt_client"

	// Topic to subscribe and publish to
	topic := "direct/aj"
	// topic := "test/topic"

	// MQTT client options
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetUsername("user")
	opts.SetPassword("password")
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	})

	// Create and start an MQTT client
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Error connecting to MQTT broker:", token.Error())
		os.Exit(1)
	}
	defer client.Disconnect(250)

	// Subscribe to the topic
	if token := client.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println("Error subscribing to topic:", token.Error())
		os.Exit(1)
	}

	// Publish messages to the topic
	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := client.Publish(topic, 1, false, text)
		token.Wait()
		time.Sleep(1 * time.Second)
	}

	// Wait to receive messages
	time.Sleep(10 * time.Second)

	// Unsubscribe from the topic
	if token := client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		fmt.Println("Error unsubscribing from topic:", token.Error())
		os.Exit(1)
	}
}
