package main

import (
	"fmt"
	"time"
)

func main() {

	bc_timer := time.NewTicker(time.Duration(10) * time.Second)

	for {
		select {
		case <-bc_timer.C:
			fmt.Println("im here ")
		}
	}

}
