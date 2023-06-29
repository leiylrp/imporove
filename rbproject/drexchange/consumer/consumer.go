package main

import (
	"github.com/streadway/amqp"
	"log"
)
func FailOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s:%s", msg, err)
	}
}
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to Connected to RabbitMQ")
	defer conn.Close()

	//定义队列的名称
	queueNames := []string{"direct_Queue1", "direct_Queue2", "direct_Queue3", "direct_Queue4"}

	//获取一个通道
	ch, err := conn.Channel()
	FailOnError(err, "Failed to Create a channel")
	defer ch.Close()

	//消费消息
	forever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	msgs, err := ch.Consume(queueNames[0], "", true, false, false, false, nil)
	FailOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message From %s : %s", queueNames[0], d.Body)
		}
	}()

	msgs, err = ch.Consume(queueNames[1], "", true, false, false, false, nil)
	FailOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message From %s : %s", queueNames[1], d.Body)
		}
	}()

	msgs, err = ch.Consume(queueNames[2], "", true, false, false, false, nil)
	FailOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message From %s : %s", queueNames[2], d.Body)
		}
	}()

	msgs, err = ch.Consume(queueNames[3], "", true, false, false, false, nil)
	FailOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message From %s : %s", queueNames[3], d.Body)
		}
	}()
	<-forever
}