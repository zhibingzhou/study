package main

import (
	"fmt"
	"net/http"
	"rabbitmq/controller"
	"rabbitmq/rabbitsimple"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	rabbitsimple.Viper()

	r := rabbitsimple.NewDirect("exchangeDirect", "keycba")
	go r.DirectConsume()

	var Router = gin.Default()
	ro := Router.Group("auth")
	ro.POST("direct", controller.Direct)
	address := fmt.Sprintf(":%s", "8088")

	s := initServer(address, Router)
	s.ListenAndServe()
}

type server interface {
	ListenAndServe() error
}

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
