package main

import (
	"fmt"
	"time"
)

type WorkerPool2 struct {
	lens    int
	Req     chan Request
	Worreq  chan chan Request
	PoolRes chan Result
}

type Request interface {
	DOsomething() Result
}

type Result struct {
	abc int
}

func NewWorkersPool(lens int) *WorkerPool2 {
	return &WorkerPool2{
		lens:    lens,
		Req:     make(chan Request),
		Worreq:  make(chan chan Request),
		PoolRes: make(chan Result),
	}
}

func (w *WorkerPool2) Run() {

	for i := 0; i < w.lens; i++ {
		mm := newwhatever()
		mm.WorkRun(w.Worreq, w.PoolRes)
	}

	go func() {
		var queReq []Request
		var queWorreq []chan Request
		for {
			var Req1 Request
			var Worreq1 chan Request
			if len(queReq) > 0 && len(queWorreq) > 0 {
				Req1 = queReq[0]
				Worreq1 = queWorreq[0]
				fmt.Println(len(queReq), len(queWorreq))
			}
			select {
			case w := <-w.Worreq:
				queWorreq = append(queWorreq, w)
			case r := <-w.Req:
				queReq = append(queReq, r)
			case Worreq1 <- Req1:
				queReq = queReq[1:]
				queWorreq = queWorreq[1:]
			}
		}
	}()

	// go func() {

	// }()
}

func (w whatever) WorkRun(in chan chan Request, PoolRes chan Result) {

	go func() {
		for {
			in <- w.R
			re := <-w.R
			result := re.DOsomething()
			//处理
			PoolRes <- result
		}
	}()

}

type whatever struct {
	R chan Request
}

func newwhatever() whatever {
	return whatever{R: make(chan Request)}
}

func (j Jm) DOsomething() Result {
	time.Sleep(time.Second * 10)
	return Result{abc: j.abc}

}

type Jm struct {
	abc     int
	PoolRes Result
}
