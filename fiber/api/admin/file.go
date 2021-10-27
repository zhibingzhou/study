package admin

import (
	"github.com/gofiber/fiber/v2"
)

type FileHistory struct {
}

func (d *FileHistory) HistoryDel(c *fiber.Ctx) error {
	filename := c.FormValue("filename")
	password := c.FormValue("password")

	return adminService.HistoryDel(filename, password, c)

}
