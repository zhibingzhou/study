package main

// type wokerall struct {
// 	lens int
// 	Rq   chan Re
// 	Rqf  chan chan Re
// 	Gv   chan Ge
// }

// func NEwwokerall(i int) *wokerall {
// 	return &wokerall{
// 		lens: i,
// 		Rq:   make(chan Re),
// 		Rqf:  make(chan chan Re),
// 		Gv:   make(chan Ge),
// 	}
// }

// func (w *wokerall) Run() {

// 	for i := 0; i < w.lens; i++ {
// 		r := NewReT()
// 		r.Gorun(w.Rqf, w.Gv)
// 	}

// 	go func() {
// 		for {
// 			select {
// 			case r := <-w.Rq:
// 				rqf := <-w.Rqf
// 				rqf <- r
// 			}
// 		}
// 	}()

// }

// func NewReT() Re {
// 	return Re{r: Re}
// }

// func (r ReT) Gorun(re chan chan Re, gv chan Ge) {
// 	go func() {
// 		for {
// 			re <- r.r
// 			select {
// 			case w := <-r.r:
// 				w.DOsomething()
// 				gv <- w.Back()
// 			}

// 		}
// 	}()
// }

// type ReT struct {
// 	r chan Re
// }

// type Re interface {
// 	DOsomething()
// 	Back() Ge
// }

// type Ge struct {
// }
