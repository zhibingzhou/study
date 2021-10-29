package admin

import (
	"studyfiber/middleware"
	"studyfiber/model/response"

	"github.com/gofiber/fiber/v2"
)

type FileHistory struct {
}

func (d *FileHistory) HistoryDel(c *fiber.Ctx) error {
	filename := c.FormValue("filename")
	password := c.FormValue("password")

	if !middleware.CheckPamarm(filename, password) {
		return response.FailWithMessage("有特殊字符存在！！", c)
	}
	return adminService.HistoryDel(filename, password, c)

}

func (f *FileHistory) CheckFileName(c *fiber.Ctx) error {
	filename := c.FormValue("filename")

	if filename == "" {
		return response.FailWithMessage("加密字串不能为空！！", c)
	}
	if !middleware.CheckPamarm(filename) {
		return response.FailWithMessage("有特殊字符存在！！", c)
	}

	return adminService.CheckFileName(filename, c)
}
