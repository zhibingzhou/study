package test

import (
	"fmt"
	"testing"
	"time"
)

/*两个协程交替打印1-100的数字 */
func TestShowNumber(t *testing.T) {
	num := make(chan int)
	go one(num)
	go two(num)

	time.Sleep(5 * time.Second)
}

func one(num chan int) {
	for i := 1; i <= 100; i++ {
		num <- i
		if i%2 == 1 {
			fmt.Println(i)
		}

	}
}

func two(num chan int) {

	for i := 1; i <= 100; i++ {
		<-num
		if i%2 == 0 {
			fmt.Println(i)
		}

	}
}
