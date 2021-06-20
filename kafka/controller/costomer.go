package controller

import (
	"kafka/thread"

	"github.com/gin-gonic/gin"
)

func StartCostomer(c *gin.Context) {
	err := thread.StartCostomer()
	if err != nil {
		return
	}
}
