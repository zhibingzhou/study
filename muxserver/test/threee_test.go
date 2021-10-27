package test

import (
	"fmt"
	"sync"
	"testing"
)

// 互斥锁的实现  mux or  chan

type AllCheck interface {
	Add()
}

func TestHc(t *testing.T) {
	chmux := CheckMux{}
	for i := 0; i < 100; i++ {
		chmux.wg.Add(1)
		go chmux.Add()
	}
	chmux.wg.Wait()
	fmt.Println(chmux.num)

	anmux := CheckChan{}
	anmux.cc = make(chan int, 1)

	for i := 100; i < 200; i++ {
		anmux.wg.Add(1)
		go anmux.Add()
	}
	anmux.wg.Wait()
	fmt.Println(anmux.num)

}

type CheckMux struct {
	mtx sync.Mutex
	wg  sync.WaitGroup
	num int
}

func (c *CheckMux) Add() {
	c.mtx.Lock()
	c.num++
	c.mtx.Unlock()
	c.wg.Done()
}

type CheckChan struct {
	wg  sync.WaitGroup
	num int
	cc  chan int
}

func (c *CheckChan) Add() {
	c.cc <- 1
	c.num++
	<-c.cc
	c.wg.Done()
}

func CheckMuxtwo() {
	var sy sync.WaitGroup

	for i := 0; i < 10; i++ {
       sy.Add(1)
	   go func(scy sync.WaitGroup){
             scy.Done()
	   }(sy)
	}
	sy.Wait()
	fmt.Println("game over !!!!!")
}
