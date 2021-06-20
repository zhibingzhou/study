package controller

import (
	"kafka/thread"

	"github.com/gin-gonic/gin"
)

func StartProducer(c *gin.Context) {
	err := thread.StartProducer()
	if err != nil {
		return
	}
}
