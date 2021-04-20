package rabbitsimple

import (
	"fmt"

	"github.com/streadway/amqp"

)

func NewFanout(exchange string) *RabbitMQ {
	c := NewRabbitMQ("", exchange, "")
	return c
}

func (r *RabbitMQ) PulishFanout(message string) {
	//name, kind string, durable, autoDelete, internal, noWait bool, args Table
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//exchange, key string, mandatory, immediate bool, msg Publishing
	err = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			Body:        []byte(message),
			ContentType: "ContentType",
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Fanout Send Message :" + message)
}

func (r *RabbitMQ) ConsumeFanout() {

	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	// name string, durable, autoDelete, exclusive, noWait bool, args Table
	q, err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)

	//name, key, exchange string, noWait bool, args Table
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}
	//queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args Table
	result, err := r.channel.Consume(
		r.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	forver := make(chan bool)

	go func() {
		for b := range result {
			fmt.Printf("Fanout Receive Message : %s", b.Body)
			fmt.Println()
		}
	}()

	<-forver

}
