package controller

import (
	"rabbitmq/rabbitsimple"

	"github.com/gin-gonic/gin"
)

func Direct(c *gin.Context) {
	r := rabbitsimple.NewDirect("exchangeDirect", "keycba")
	go r.DirectConsume()
	return 
}
