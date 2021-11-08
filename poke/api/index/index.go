package index

import (
	"github.com/astaxie/beego/logs"
     "github.com/gofiber/fiber/v2"
	"text/template"
)

type IndexPage struct {

}

func (i *IndexPage) Index(c *fiber.Ctx) {

	_,err := template.ParseFiles("templates/poker.html")
	if err != nil {
		logs.Error("user request Index - can't find template file %s", "templates/poker.html")
		return
	}


}