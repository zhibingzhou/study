package router

import (
	"kafka/controller"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	Router = gin.New()
	//静态文件路径，一定需要
	// Router.LoadHTMLGlob("view/*")
	// Router.LoadHTMLFiles("./view/index.html", "./view/city.html", "./view/user.html", "./view/work.html", "./view/province.html", "./view/indexvue.html", "./view/videoindex.html")
	// Router.Static("/layui", "./layui")
	// Router.GET("/costomer", controller.Home)
	// Router.GET("/video", controller.Video)

	//API接口 消费者
	RuoAi := Router.Group("/costomer")
	//用户
	RuoAi.POST("/start", controller.StartCostomer)

	//API接口 生产者
	XinHe := Router.Group("/producer")
	XinHe.POST("/start", controller.StartProducer)

}
