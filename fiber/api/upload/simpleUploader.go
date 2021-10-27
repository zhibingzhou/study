package upload

import (
	"fmt"
	"studyfiber/model"
	"studyfiber/model/response"
	"studyfiber/utils"

	"github.com/gofiber/fiber/v2"
)

type SimpleUplader struct {
}

// @Tags SimpleUploader
// @Summary 断点续传插件版示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "断点续传插件版示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"切片创建成功"}"
// @Router /SimpleUploaderApi/upload [post]
func (s *SimpleUplader) SimpleUploaderUpload(c *fiber.Ctx) error {
	var chunk model.ExaSimpleUploader
	header, err := c.FormFile("file")
	chunk.Filename = c.FormValue("filename")
	chunk.ChunkNumber = c.FormValue("chunkNumber")
	chunk.CurrentChunkSize = c.FormValue("currentChunkSize")
	chunk.Identifier = c.FormValue("identifier")
	chunk.TotalSize = c.FormValue("totalSize")
	chunk.TotalChunks = c.FormValue("totalChunks")
	var chunkDir = "./chunk/" + chunk.Identifier + "/"
	hasDir, _ := utils.PathExists(chunkDir)
	if !hasDir {
		if err := utils.CreateDir(chunkDir); err != nil {
			//	global.GVA_LOG.Error("创建目录失败!", zap.Any("err", err))
		}
	}
	chunkPath := chunkDir + chunk.Filename + chunk.ChunkNumber
	err = c.SaveFile(header, chunkPath)
	if err != nil {
		// global.GVA_LOG.Error("切片创建失败!", zap.Any("err", err))
		return response.FailWithMessage("切片创建失败", c)

	}
	chunk.CurrentChunkPath = chunkPath

	//数据库写入
	err = uploadService.SaveChunk(chunk)
	if err != nil {
		//global.GVA_LOG.Error("切片创建失败!", zap.Any("err", err))
		return response.FailWithMessage("切片创建失败", c)

	} else {
		return response.OkWithMessage("切片创建成功", c)
	}
}

// @Tags SimpleUploader
// @Summary 断点续传插件版示例
// @Security ApiKeyAuth
// @Produce  application/json
// @Param md5 query string true "md5"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /SimpleUploaderApi/checkFileMd5 [get]
func (s *SimpleUplader) CheckFileMd5(c *fiber.Ctx) error {
	md5 := c.Query("md5")
	err, chunks, isDone := uploadService.CheckFileMd5(md5)
	if err != nil {
		// global.GVA_LOG.Error("md5读取失败!", zap.Any("err", err))
		return response.FailWithMessage("md5读取失败", c)
	} else {
		return response.OkWithDetailed(fiber.Map{
			"chunks": chunks,
			"isDone": isDone,
		}, "查询成功", c)
	}
}

// @Tags SimpleUploader
// @Summary 合并文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param md5 query string true "md5"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"合并成功"}"
// @Router /SimpleUploaderApi/mergeFileMd5 [get]
func (s *SimpleUplader) MergeFileMd5(c *fiber.Ctx) error {
	md5 := c.Query("md5")
	fileName := c.Query("fileName")

	//aes := common.SetAES(utils.AppConf.PrivateKey, utils.AppConf.Iv, "", 32)
	//fileName = aes.AesDecryptString(fileName)

	//fileName, _ = url.QueryUnescape(fileName)
	fmt.Println(fileName)
	err := uploadService.MergeFileMd5(md5, fileName)
	if err != nil {
		// global.GVA_LOG.Error("md5读取失败!", zap.Any("err", err))
		return response.FailWithMessage("md5读取失败", c)
	} else {
		return response.OkWithMessage("合并成功", c)
	}
}
