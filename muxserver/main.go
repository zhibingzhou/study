package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"webstudy/controller"
	"webstudy/global"
	"webstudy/session"
)

type test struct {
	aaa bool
}

func main() {
	global.Global.Sessionstore = session.NewSessionStore(15 * time.Second)
	fmt.Println(test{})
	mux := http.NewServeMux()
	rh := &controller.Testhander{Format: time.RFC1123}
	mux.Handle("/time", rh)

	//正常这个就够用
	mux.HandleFunc("/stime", controller.SecondTest)

	//需要传参的可以用这个
	mux.Handle("/three/123", controller.ThreeTest(time.RFC3339Nano))
	//静态目录
	mux.Handle("/tmp/abc/", http.StripPrefix("/tmp/abc/", http.FileServer(http.Dir(`D:\zhou\src\study\muxserver\tmp`))))

	//websocket 连接
	mux.HandleFunc("/webchannel", controller.WebSocketConnect)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
	fmt.Println("ehree")
}
