package pkg

import (
	"context"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Publisher struct {
	Topic    string
	QoS      byte
	Retained bool
	Client   mqtt.Client
}

func (p *Publisher) Publish(ctx context.Context, msg any) error {
	token := p.Client.Publish(p.Topic, p.QoS, p.Retained, msg)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	fmt.Printf("publish to :%s, msg %s\n", p.Topic, msg)
	return nil
}
