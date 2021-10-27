package test

import "github.com/gofiber/fiber/v2"

/*测试首页*/
type IndexApi struct {
}

func (i *IndexApi) Test(c *fiber.Ctx) error {
	host := c.Request().Host()
	return c.Render("view/index.tpl", fiber.Map{
		"title": "测试接口",
		"host":  string(host),
	})
}
