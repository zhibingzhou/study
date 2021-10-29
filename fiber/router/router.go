package router

import (
	"studyfiber/api"
	"studyfiber/middleware"

	"github.com/gofiber/fiber/v2"
)

var Router *fiber.App

func init() {

	Router = fiber.New()
	var test_Hellorouter = api.ApiGroupApp.TestApiGroup.HelloWorldApi
	// 路由：/ hello world
	index := Router.Group("/").Use(middleware.Check)
	index.Get("/", test_Hellorouter.World)

	AdminRouterInit()
	TestRouterInit()
	UploadRouterInit()

}
