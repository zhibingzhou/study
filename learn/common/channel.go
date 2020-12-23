package common

import (
	"fmt"

	"github.com/axgle/mahonia"
)

type ChannelPool struct {
	lens    int
	WorkGO  chan Request
	GworkGo chan chan Request
	Result  chan ParseResult
	QorS    bool
}

func NewChannelPool(i int, q bool) *ChannelPool {
	return &ChannelPool{
		lens: i,
		WorkGO:  make(chan Request),
		GworkGo: make(chan chan Request),
		Result:  make(chan ParseResult),
		QorS:    q, //是否需要有序队列
	}
}

func (s *ChannelPool) Run() {

	for i := 0; i < s.lens; i++ {
		r := NewWokerS()
		r.Run(s.GworkGo, s.Result)
	}

	//两种通道跑方式  管理
	if s.QorS == false {
		go func() {
			for {
				select {
				case work := <-s.WorkGO:
					gwork := <-s.GworkGo
					gwork <- work
					// fmt.Println("我是随机")
				}
			}
		}()
	} else {

		go func() {
			var qWorkGO []chan Request
			var qWork []Request
			for {
				var workgo chan Request
				var work Request
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
					// fmt.Println("我是队列") 有序，先进先出
					qWorkGO = qWorkGO[1:]
					qWork = qWork[1:]
				}
			}
		}()
	}
}

type WokerS struct {
	c chan Request
}


//工作的地方
func (w WokerS) Run(gr chan chan Request, r chan ParseResult) {

	go func() {
		for {
			gr <- w.c
			select {
			case m := <-w.c:
				result, _ := m.worker()
				r <- result
			}
		}

	}()

}

func NewWokerS() WokerS {
	return WokerS{c: make(chan Request)}
}

func (r Request) worker() (ParseResult, error) {

	body, err := Fetch(r.Url)

	bodystr := mahonia.NewDecoder("gbk").ConvertString(string(body))

	if err != nil {
		fmt.Println("Fetch:", err.Error())
		return ParseResult{}, err
	}
	parseResult := r.ParserFunc([]byte(bodystr))

	return parseResult, nil
}
