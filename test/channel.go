package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

//定义结构
type WorkerPool struct {
	workerlen   int
	JobQueue    chan Job
	WorkerQueue chan chan Job
	PoolRes     chan PoolResult
}

//定义接口 Job
type Job interface {
	Do()
	GetResult() PoolResult
}

//定义结构
type PoolResult struct {
	Status  int
	Msg     string
	JsonRes string
}

//实例化 WorkerPool  1.
func NewWorkerPool(workerlen int) *WorkerPool {
	return &WorkerPool{
		workerlen:   workerlen,
		JobQueue:    make(chan Job),
		WorkerQueue: make(chan chan Job, workerlen),
		PoolRes:     make(chan PoolResult, workerlen),
	}
}

func (wp *WorkerPool) Run() {
	//初始化worker
	for i := 0; i < wp.workerlen; i++ {
		worker := NewWorker() //实例化一个worker
		worker.Run(wp.WorkerQueue, wp.PoolRes)
	}
	// 循环获取可用的worker,往worker中写job
	go func() {
		for {
			select {
			case job := <-wp.JobQueue: // ==> chan job 拿出 job
				fmt.Println("im here one")
				worker := <-wp.WorkerQueue // chan chan Job 拿出  chan job
				worker <- job              //把 job 塞到 chan job  job -> chan job
			}
		}
	}()
}

type Worker struct {
	JobQueue chan Job
}

//实例化一个worker
func NewWorker() Worker {
	return Worker{JobQueue: make(chan Job)}
}

//Worker Run 3.
func (w Worker) Run(wq chan chan Job, ps chan PoolResult) {
	go func() {
		for {
			wq <- w.JobQueue //等待传送 chan job 到  chan chan Job  说明这个worker 在等待了
			fmt.Println("im here two")
			select {
			case job := <-w.JobQueue: //等待传送 job 到  Job
				job.Do()
				ps <- job.GetResult()
			}
		}
	}()
}

//定义一个 UpdateCash 结构来实现Job接口
type UpdateCash struct {
	Order_number string
	Pay_order    string
	Note         string
	Order_type   int
	PoolRes      PoolResult
}

//实现接口 Job
func (uc *UpdateCash) Do() {
	// t_status := 100
	// t_msg := "额度更新错误"
	// //order_type:1=支付,2=下发,3=代付失败返回,4=后台下发到卡,5=上游下发失败返还,6=纯代付下发
	// switch uc.Order_type {
	// case 1:
	// 	t_status, t_msg = updatePay(uc.Order_number, uc.Pay_order, uc.Note)
	// case 2:
	// 	t_status, t_msg = cashOrder(uc.Order_number, uc.Pay_order, uc.Note)
	// case 3:
	// 	t_status, t_msg = updateReturnCash(uc.Order_number, uc.Pay_order, uc.Note)
	// case 4:
	// 	t_status, t_msg = order(uc.Order_number, uc.Pay_order, uc.Note)
	// case 5:
	// 	t_status, t_msg = updateReturnOrder(uc.Order_number, uc.Pay_order, uc.Note)
	// case 6:
	// 	t_status, t_msg = dfOrder(uc.Order_number, uc.Pay_order, uc.Note)
	// }
	time.Sleep(time.Second * 2)
	fmt.Println("im Do")
	uc.PoolRes = PoolResult{Status: 1, Msg: "123", JsonRes: ""}

}

func (uc *UpdateCash) GetResult() PoolResult {
	return uc.PoolRes
}

type OrderList struct {
	Id           string `orm:"pk"`
	Status       int    //订单状态(1=处理中,3=完成,9=拒绝,-1=未扣款的废单)
	Pay_code     string
	Pay_id       int
	Amount       float64
	Real_amount  float64
	Create_time  string
	Pay_time     string
	Order_number string
	Cash_id      string
	Pay_order    string
	Bank_code    string
	Note         string
	Bank_title   string
	Branch       string
	Card_name    string
	Card_number  string
	Phone        string
	Fee_amount   float64
	Order_amount float64
	Order_type   int //是否纯代付下发:1=代收下发,2=纯代付下发
}

type MerList struct {
	Id          int
	Code        string
	Domain      string
	Title       string
	Qq          string
	Skype       string
	Telegram    string
	Phone       string
	Email       string
	Private_key string
	Is_agent    int
	Agent_path  string
	Amount      float64
	Total_in    float64
	Total_out   float64
}

type PayList struct {
	Id           string `orm:"pk"`
	Status       int    //订单状态(1=处理中,3=完成,9=拒绝)
	Pay_code     string
	Pay_id       int
	Mer_code     string
	Push_status  int
	Push_num     int
	Amount       float64
	Real_amount  float64
	Create_time  string
	Pay_time     string
	Order_number string
	Pay_order    string
	Class_code   string
	Bank_code    string
	Push_url     string
	Note         string
	Is_mobile    int
	Rate         float64
	Agent_path   string
}

type PayData struct {
	Amount       string //订单金额
	Order_number string //订单编号
	Class_code   string //支付类型
	Pay_bank     string //选择的银行或第三方支付平台
	Is_mobile    string //是否手机版 0是网页  1是手机
	Ip           string //客户的IP
}

type Gorm struct {
	DB *gorm.DB
}

type CountTotal struct {
	Total float64
	Num   int
}

type MerRate struct {
	Id           int
	Mer_code     string
	Pay_code     string
	Class_code   string
	Bank_code    string
	Rate         float64
	Limit_amount float64
	Day_amount   float64
}

var gdb Gorm
