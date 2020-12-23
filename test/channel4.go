package main

import (
	"fmt"
)

type StudentPool struct {
	lens    int
	WorkGO  chan Change
	GworkGo chan chan Change
	Result  chan ResultStatus
	QorS    bool
}

type Change interface {
	ChangeScore() ResultStatus
	ChangeGander() ResultStatus
}

func NewStudentPool(i int, q bool) *StudentPool {
	return &StudentPool{lens: i,
		WorkGO:  make(chan Change),
		GworkGo: make(chan chan Change),
		Result:  make(chan ResultStatus),
		QorS:    q,
	}
}

func (s *StudentPool) Run() {

	for i := 0; i < s.lens; i++ {
		r := NewWokerS()
		r.Run(s.GworkGo, s.Result)
	}
	if s.QorS == false {
		go func() {
			for {
				select {
				case work := <-s.WorkGO:
					gwork := <-s.GworkGo
					gwork <- work
					fmt.Println("我是随机")
				}
			}
		}()
	} else {

		go func() {
			var qWorkGO []chan Change
			var qWork []Change
			for {
				var workgo chan Change
				var work Change
				if len(qWorkGO) > 0 && len(qWork) > 0 {
					workgo = qWorkGO[0]
					work = qWork[0]
				}
				select {
				case qw := <-s.GworkGo:
					qWorkGO = append(qWorkGO, qw)
				case w := <-s.WorkGO:
					qWork = append(qWork, w)
				case workgo <- work:
					fmt.Println("我是队列")
					qWorkGO = qWorkGO[1:]
					qWork = qWork[1:]
				}
			}
		}()
	}
}

type WokerS struct {
	c chan Change
}

func (w WokerS) Run(gr chan chan Change, r chan ResultStatus) {

	go func() {
		for {
			gr <- w.c
			select {
			case m := <-w.c:
				result := m.ChangeScore()
				fmt.Println("换分数", result)
				result = m.ChangeGander()
				fmt.Println("换性别", result)
				r <- result
			}
		}

	}()

}

func NewWokerS() WokerS {
	return WokerS{c: make(chan Change)}
}

type ResultStatus struct {
	status int
	msg    string
}

type GotoChange struct {
	Score  int
	Gander string
	Change bool
}

func (g GotoChange) ChangeScore() ResultStatus {
	return ResultStatus{status: 200, msg: "分数"}
}

func (g GotoChange) ChangeGander() ResultStatus {
	return ResultStatus{status: 200, msg: "性别"}
}
