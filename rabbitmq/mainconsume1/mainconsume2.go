package main

import "rabbitmq/rabbitsimple"

import "fmt"

func main() {

	// r := rabbitsimple.NewRabbitMQSimple("nimabi")
	// 	r.ConsumeSimple()

	// r := rabbitsimple.NewDirect("exchangeDirect", "keycba")

	// r.DirectConsume()

	r := rabbitsimple.NewTopic("exchangeTopic", "*keycba*")
	r.TopicConsume()
	
}

func woker(id int, c <-chan int) {
	for r := range c {
		fmt.Printf("Id = %d, chan = %c\n", id, r)
	}
}
