package main

import (
	_ "studyfiber/model"
	"studyfiber/router"
	_ "studyfiber/utils"
)

/**/

func main() {

	// 在 8001 端口启动服务
	router.Router.Listen(":8001")
}
