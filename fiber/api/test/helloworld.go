package test

import "github.com/gofiber/fiber/v2"

/*测试首页*/
type HelloWorldApi struct {
}

func (i *HelloWorldApi) World(c *fiber.Ctx) error {
	return c.SendString("Hi! its from fiber !! Have a good day!")
}
