package main

import (
	"context"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/egigiffari/broker-go-mqtt/pkg"
)

func main() {
	client := getClient()
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	pubs := getPublisher(client)

	for i := 1; i <= 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := pubs.Publish(ctx, fmt.Sprintf("message %d", i))
		if err != nil {
			fmt.Printf("failed publish, err: %s", err.Error())
		}
	}

	client.Disconnect(500)
}

func getClient() mqtt.Client {
	config := pkg.Config{
		Host:     "localhost",
		Port:     1883,
		ClientId: "1234",
		Username: "",
		Password: "",
	}

	return pkg.NewClient(&config)
}

func getPublisher(client mqtt.Client) *pkg.Publisher {
	return &pkg.Publisher{
		Topic:    "topic/test",
		QoS:      2,
		Retained: true,
		Client:   client,
	}
}
