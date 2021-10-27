package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"webstudy/session"
	"webstudy/store/types"
	"webstudy/thread"

	"github.com/tinode/chat/server/logs"
)

type Testhander struct {
	Format  string
	Request string
	Method  string
}

var adc chan int

//定死这样写  Handler 里面有这个接口目的实现他
func (t *Testhander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(t.Format)
	w.Write([]byte("The time is: " + tm)) // 返回一个字串
}

func SecondTest(w http.ResponseWriter, r *http.Request) {

	apikey := r.Header.Get("X-Tinode-APIKey") //可以拿到header 头

	if apikey == "" {
		apikey = r.URL.Query().Get("apikey") //可以拿到get请求后面的元素
	}
	if apikey == "" {
		apikey = r.FormValue("apikey") //可以拿到post 请求后面的元素
	}
	if apikey == "" {
		co, _ := r.Cookie("apikey") //可以拿到cookie 里面的元素
		apikey = co.Value
	}

	// tm := time.Now().Format(time.RFC3339)
	// aa := strconv.Itoa(len(adc))
	// w.Write([]byte("The time is: " + tm + "len(chan)" + aa))

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Testhander{Format: "测试返回一个json", Request: apikey, Method: r.Method})
}

func ThreeTest(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		abc := []string{"1", "2", "3"}
		fmt.Println(abc[1:])
		fmt.Println(abc[:2])
		tm := time.Now().Format(format)
		go func() {
			adc <- 1
		}()
		w.Write([]byte("The time is: " + tm))
	}
	return http.HandlerFunc(fn)
}

func WebSocketConnect(w http.ResponseWriter, r *http.Request) {

	// if isValid, _ := checkAPIKey(getAPIKey(req)); !isValid {
	// 	wrt.WriteHeader(http.StatusForbidden)
	// 	json.NewEncoder(wrt).Encode(ErrAPIKeyRequired(now))
	// 	logs.Err.Println("ws: Missing, invalid or expired API key")
	// 	return
	// }
	apikey := r.Header.Get("X-Tinode-APIKey") //可以拿到header 头

	if apikey == "" {
		apikey = r.URL.Query().Get("apikey") //可以拿到get请求后面的元素
	}
	if apikey == "" {
		apikey = r.FormValue("apikey") //可以拿到post 请求后面的元素
	}
	if apikey == "" {
		co, err := r.Cookie("apikey") //可以拿到cookie 里面的元素
		if err == nil {
			apikey = co.Value
		}

	}

	if apikey == "" {

	}
	now := types.TimeNow()
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(session.ErrOperationNotAllowed("", "", now))
		logs.Err.Println("ws: Invalid HTTP method", r.Method)
		return
	}

	//将HTTP协议升级到Websocket协议
	thread.WebSocketConnect(w, r)
}

func init() {
	// 管道，只有定义了大小并且有输入，才会有大小，两个条件必须都满足
	adc = make(chan int, 10)
	go func() {
		for {
			select {
			case a := <-adc:
				fmt.Println(a, len(adc))
				time.Sleep(10 * time.Second)
			}
		}
	}()
	add = Same(10, Add)

	add()
}

var add func() int

func Same(id int, a func() int) func() int {
	return func() int {
		//可以做 check 动作
		a()
		return 10
	}
}

func Add() int {
	fmt.Println("ADD")
	return 10
}
