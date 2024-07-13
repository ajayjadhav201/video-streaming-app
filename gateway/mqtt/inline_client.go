package mqtt

import (
	"gateway/utils"
	"time"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

// this client
type InlineClient struct {
	Client *Mqtt
}

// this client listens to users requests and send responses to them
func NewInlineClient(mq *Mqtt) *InlineClient {
	return &InlineClient{Client: mq}
}

func (mq *Mqtt) InitInlineClient() {
	go func() {
		time.Sleep(10 * time.Second)
		utils.LogPrintln("inline client Subscribed")
		_ = mq.Subscribe("direct/#", 1, subscribeCallback)
		// _ = mq.Subscribe("direct/#", 2, subscribeCallback)
		mq.Publish("direct/aj", []byte("this is inline message"), false, 1)
	}()
}

func subscribeCallback(c *mqtt.Client, sub packets.Subscription, pkt packets.Packet) {
	utils.LogPrintln("inline client received message from subscription",
		"client",
		c.ID,
		"subscriptionId",
		sub.Identifier,
		"topic",
		pkt.TopicName,
		"payload",
		string(pkt.Payload),
	)

}
