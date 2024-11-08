package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/egigiffari/broker-go-mqtt/pkg"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	subs := getSubsciber()

	go subs.Listen(ctx)

	<-ctx.Done()
	cancel()

	subs.Close(500)
}

func getSubsciber() *pkg.Subscriber {
	config := pkg.Config{
		Host:     "localhost",
		Port:     1883,
		ClientId: "123",
		Username: "",
		Password: "",
	}

	return &pkg.Subscriber{
		Topic:  "topic/test",
		QoS:    1,
		Config: &config,
		Handler: func(c mqtt.Client, msg mqtt.Message) {
			fmt.Printf("recieved: %s\n", string(msg.Payload()))
		},
	}
}
