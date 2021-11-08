package router

import "github.com/gofiber/fiber/v2"

var Router *fiber.App

func init() {
	Router = fiber.New()
	fiber.New()
	//静态文件目录
	Router.Static("/static", "./static")
}
