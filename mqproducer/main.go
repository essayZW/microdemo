package main

import (
	"log"
	"strconv"
	"time"

	"github.com/asim/go-micro/plugins/broker/rabbitmq/v4"
	"go-micro.dev/v4/broker"
)

func main() {
	rabbitmqBroker := rabbitmq.NewBroker(
		broker.Addrs("amqp://172.17.0.3:5672"),
	)
	if err := rabbitmqBroker.Init(); err != nil {
		log.Panicln(err)
	}

	if err := rabbitmqBroker.Connect(); err != nil {
		log.Panicln(err)
	}

	ticker := time.NewTicker(time.Duration(2) * time.Second)
	id := 0
	for time := range ticker.C {
		id++
		rabbitmqBroker.Publish("micro.demo.test", &broker.Message{
			Header: map[string]string{
				"Id":   strconv.Itoa(id),
				"Time": time.GoString(),
			},
			Body: []byte("test"),
		})
		log.Printf("Publish %d", id)
	}

}
