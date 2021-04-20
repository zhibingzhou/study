package rabbitsimple

import (
	"fmt"

	"github.com/streadway/amqp"

)

func NewTopic(exchange, key string) *RabbitMQ {
	r := NewRabbitMQ("", exchange, key)
	return r
}

func (r *RabbitMQ) TopicPulish(message string) {

	//name, kind string, durable, autoDelete, internal, noWait bool, args Table
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}
	//exchange, key string, mandatory, immediate bool, msg Publishing

	r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			Body:        []byte(message),
			ContentType: "text/pain",
		},
	)
	fmt.Println("Topic send ", message)

}

func (r *RabbitMQ) TopicConsume() {

	//name, kind string, durable, autoDelete, internal, noWait bool, args Table
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	//name string, durable, autoDelete, exclusive, noWait bool, args Table
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		true,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}
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
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	forver := make(chan bool)
	for b := range result {
		fmt.Printf("Direct Receive Message : %s", b.Body)
		fmt.Println()
	}
	<-forver
}

func severityFrom(args []string) string {
	if len(args) > 1 {
		return args[1]
	}
	return "anonymous.info"
}
