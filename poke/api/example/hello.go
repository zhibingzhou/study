package example

import "github.com/gofiber/fiber/v2"

type HelloWorld struct {
}

func (h *HelloWorld) Show(f *fiber.Ctx) error {
	return f.SendString("Hello World , this is from fiber !!")
}
