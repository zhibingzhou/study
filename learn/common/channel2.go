package common

type Channel_Pool struct {
	Count    int
	Request  chan DRequest
	Requests chan chan DRequest
	Result   chan DResult
}

type DRequest struct {
	Id     int
	Type   string
	Amount int
}

type DResult struct {
	Id    int
	Solve string
}

func NewChannel_Pool(count int) *Channel_Pool {
	return &Channel_Pool{
		Count:    count,
		Request:  make(chan DRequest),
		Requests: make(chan chan DRequest),
		Result:   make(chan DResult)}
}

func (c *Channel_Pool) NewChannel_PoolGo() {

	for i := 0; i < c.Count; i++ { //工作
		w := Newwoker()
		w.Run(c.Requests, c.Result)
	}

	// go func() {
	// 	for {
	// 		select {
	// 		case woker := <-c.Request:
	// 			wokers := <-c.Requests
	// 			wokers <- woker
	// 		}
	// 	}
	// }()

	go func() {

		var DRs []chan DRequest
		var DR []DRequest

		for {

			var goRs chan DRequest
			var goDR DRequest

			if len(DRs) > 0 && len(DR) > 0 {
				goRs = DRs[0]
				goDR = DR[0]
			}

			select {

			case requests := <-c.Requests:
				DRs = append(DRs, requests)
			case request := <-c.Request:
				DR = append(DR, request)
			case goRs <- goDR:
				DRs = DRs[1:]
				DR = DR[1:]
			}

		}

	}()

}

//工作 发送管理，接收任务
//管理器，接受管理，分配任务

type Woker struct {
	Wrequest chan DRequest
}

func Newwoker() Woker {
	return Woker{Wrequest: make(chan DRequest)}
}

//工作
func (w Woker) Run(crequest chan chan DRequest, cresult chan DResult) {

	go func() {

		for {
			crequest <- w.Wrequest
			select {
			case r := <-w.Wrequest:
				result := r.Run()
				cresult <- result
			}
		}

	}()
}

func (d DRequest) Run() DResult {
	return DResult{Solve: d.Type, Id: d.Id}
}
