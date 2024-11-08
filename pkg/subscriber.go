package pkg

import (
	"context"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Subscriber struct {
	Topic   string
	QoS     byte
	Handler mqtt.MessageHandler
	Config  *Config
	client  mqtt.Client
}

func (s *Subscriber) Listen(ctx context.Context) {
	opts := s.Config.Get()
	opts.SetOnConnectHandler(s.subscribe)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		fmt.Printf("failed to connect, err: %s\n", token.Error().Error())
		return
	}

	s.client = client
	<-ctx.Done()
}

func (s *Subscriber) Close(quiesce uint) {
	s.client.Disconnect(quiesce)
}

func (s *Subscriber) subscribe(c mqtt.Client) {
	token := c.Subscribe(s.Topic, s.QoS, s.Handler)
	if token.Wait() && token.Error() != nil {
		log.Printf("token err : %v", token.Error())
		return
	}

	fmt.Printf("subscribe :%s\n", s.Topic)
}
