package main

import (
	"rabbitmq/rabbitsimple"
	"strconv"
	"time"

)

func main() {

	// r := rabbitsimple.NewRabbitMQSimple("nimabi")
	// for i := 0; i < 10; i++ {
	// 	time.Sleep(time.Second)
	// 	r.PulishSimple("ID = " + strconv.Itoa(i))
	// }

	// r := rabbitsimple.NewFanout("exchangeFanout")
	// for i := 0; i < 10; i++ {
	// 	time.Sleep(time.Second)
	// 	r.PulishFanout("ID = " + strconv.Itoa(i))
	// }

	// r := rabbitsimple.NewDirect("exchangeDirect", "keyabc")
	// for i := 0; i < 10; i++ {
	// 	time.Sleep(time.Second)
	// 	r.DirectPulish("ID = " + strconv.Itoa(i))
	// }

	// rb := rabbitsimple.NewDirect("exchangeDirect", "keycba")
	// for i := 0; i < 10; i++ {
	// 	time.Sleep(time.Second)
	// 	rb.DirectPulish("ID = " + strconv.Itoa(i))
	// }

	rb := rabbitsimple.NewTopic("exchangeTopic", `keycba.keyabc.abc`)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		rb.TopicPulish("ID = " + strconv.Itoa(i))
	}

	// r := rabbitsimple.NewTopic("exchangeTopic", "keyabc")
	// for i := 0; i < 10; i++ {
	// 	time.Sleep(time.Second)
	// 	r.TopicPulish("ID = " + strconv.Itoa(i))
	// }

}
