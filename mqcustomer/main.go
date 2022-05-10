package main

import (
	"fmt"
	"log"

	"github.com/asim/go-micro/plugins/broker/rabbitmq/v4"
	"go-micro.dev/v4/broker"
)

func main() {

	rabbitmqBroker := rabbitmq.NewBroker(
		broker.Addrs("amqp://172.17.0.4:5672"),
		rabbitmq.DurableExchange(),
	)
	if err := rabbitmqBroker.Init(); err != nil {
		log.Panicln(err)
	}

	if err := rabbitmqBroker.Connect(); err != nil {
		log.Panicln(err)
	}

	rabbitmqBroker.Subscribe("micro.demo.test", func(p broker.Event) error {
		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	}, broker.Queue("xss"), rabbitmq.AckOnSuccess())
	c := make(chan struct{})
	<-c
}
