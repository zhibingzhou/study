package main

import (
	"fmt"
	"kafka/router"
	"kafka/thread"
)

func main() {
	fmt.Println("<---start--->")
	go thread.StartCostomer()
	router.Router.Run(":8083")
}
