package test

import (
	"fmt"
	"testing"
)

/*
 启动一个协程，放入1-2000 数据，
*/

func TestRun(T *testing.T) {
	numChan := make(chan int, 2000)
	resChan := make(chan []int)
	go Send(numChan)
	for i := 0; i < 8; i++ {
		go Read(numChan, resChan)
	}
	Show(resChan)
}

func Send(numChan chan int) {
	for i := 1; i < 20; i++ {
		numChan <- i
	}
	close(numChan)
}

func Read(numChan chan int, resChan chan []int) {
	for {
		select {
		case n, ok := <-numChan:
			if !ok {
				return
			}
			if n == 0 {
				continue
			}
			sum := 0
			for i := 1; i < n; i++ {
				sum += i
			}
			all := []int{n, sum}
			resChan <- all
		}
	}
}

func Show(resChan chan []int) {
	for {
		select {
		case rs, ok := <-resChan:
			if !ok {
				break
			}
			fmt.Println(rs[0], rs[1])
		default:
			break
		}
	}
}

