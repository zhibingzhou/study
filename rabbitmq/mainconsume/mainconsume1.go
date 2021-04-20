package main

import (
	"fmt"
	"rabbitmq/rabbitsimple"

)

func main() {

	// r := rabbitsimple.NewRabbitMQSimple("nimabi")
	// 	r.ConsumeSimple()

	// r := rabbitsimple.NewFanout("exchangeFanout")
	// r.ConsumeFanout()

	r := rabbitsimple.NewTopic("exchangeTopic", `keycba.#`)

	r.TopicConsume()

}

func woker(id int, c <-chan int) {
	for r := range c {
		fmt.Printf("Id = %d, chan = %c\n", id, r)
	}
}
