package router

import (
	"learn/controller"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.New()
	//静态文件路径，一定需要
	Router.LoadHTMLGlob("view/*")
	Router.LoadHTMLFiles("./view/show.tpl", "./view/jump.tpl")
	Router.GET("/all_data", controller.GetAllData)
	Router.POST("/get_url", controller.KeepUrl)
	Router.GET("/", controller.GoLongUrl)

}
