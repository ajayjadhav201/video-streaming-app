package mqtt

import (
	"errors"
	"gateway/utils"

	mqtt "github.com/mochi-mqtt/server/v2"

	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
)

type Mqtt struct {
	Address string
	server  *mqtt.Server
}

func NewMqttServer(address string) *Mqtt {
	//init mqtt
	server := initMqtt(address)
	//
	return &Mqtt{
		Address: address,
		server:  server,
	}
}

func (mq *Mqtt) Run() error {
	if mq.server == nil {
		return errors.New("mqtt server not started yet")
	}
	return mq.server.Serve()
}

func (mq *Mqtt) Subscribe(filter string, subscriptionId int, handler mqtt.InlineSubFn) error {
	return mq.server.Subscribe(filter, subscriptionId, handler)
}

func (mq *Mqtt) Publish(topic string, payload []byte, retain bool, qos byte) error {
	return mq.server.Publish(topic, payload, retain, qos)
}

func (mq *Mqtt) Close() error {
	if mq.server == nil {
		return nil
	}
	err := mq.server.Close()
	if err != nil {
		return err
	}
	mq.server = nil
	return nil
}

func authHook() *auth.Ledger {
	authRules := &auth.Ledger{
		Auth: auth.AuthRules{
			{Username: "user", Password: "password", Allow: true},
		},
		// ACL: auth.ACLRules{
		// 	{Username: "user", Filters: auth.Filters{
		// 		"user/#":    auth.ReadWrite,
		// 		"updates/#": auth.ReadWrite,
		// 	}},
		// },
	}
	return authRules
}

func initMqtt(address string) *mqtt.Server {
	//
	// start a new mqtt server
	server := mqtt.New(&mqtt.Options{InlineClient: true})
	//
	// add auth hook
	err := server.AddHook(new(auth.Hook), &auth.Options{
		Ledger: authHook(),
	})
	utils.FatalIfError(err, "Failed to start mqtt server...")
	//
	// add listener
	tcp := listeners.NewTCP(
		listeners.Config{
			ID:      "t1",
			Address: address,
		},
	)
	err = server.AddListener(tcp)
	utils.FatalIfError(err, "Failed to add tcp listener...")
	//
	//
	return server
}
