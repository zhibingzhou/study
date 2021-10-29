package admin

import (
	"studyfiber/model"
	"studyfiber/model/response"
	"studyfiber/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type FileHistory struct {
}

func (f *FileHistory) HistoryDel(filename, password string, c *fiber.Ctx) error {
	if password != utils.AppConf.PrivateKey {
		return response.FailWithMessage("密码错误", c)

	}
	if filename == "" {
		return response.FailWithMessage("文件名不能为空", c)

	}
	sim := model.ExaSimpleUploader{}

	err := model.Gdb.DB.Where("filename like ?", `%`+filename+`%`).First(&sim).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.FailWithMessage("未找到文件", c)

		}
		return response.FailWithMessage("查询错误", c)

	}

	del := `delete from exa_simple_uploader where   filename like '%` + filename + `%';`
	err = model.Query(del)

	if err != nil {
		return response.FailWithMessage(err.Error(), c)

	}

	return response.Ok(c)
}

func (f *FileHistory) CheckFileName(chunk_number string, c *fiber.Ctx) error {
	sim := model.ExaSimpleUploader{}
	err := model.Gdb.DB.Where("chunk_number = ?", chunk_number).First(&sim).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.FailWithMessage("未找到文件", c)

		}
		return response.FailWithMessage("查询错误", c)

	}

	return response.OkWithMessage(chunk_number+" ___对应文件名：  "+sim.Filename, c)
}
