package upload

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"studyfiber/model"
	"studyfiber/utils"
	"time"

	"github.com/jinzhu/gorm"
)

type SimpleUploader struct {
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SaveChunk
//@description: 保存文件切片路径
//@param: uploader model.ExaSimpleUploader
//@return: err error

func (s *SimpleUploader) SaveChunk(uploader model.ExaSimpleUploader) (err error) {
	return model.Gdb.DB.Create(uploader).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CheckFileMd5
//@description: 检查文件是否已经上传过
//@param: md5 string
//@return: err error, uploads []model.ExaSimpleUploader, isDone bool

func (s *SimpleUploader) CheckFileMd5(md5 string) (err error, uploads []model.ExaSimpleUploader, isDone bool) {
	err = model.Gdb.DB.Find(&uploads, "identifier = ? AND is_done = ?", md5, false).Error
	isDone = errors.Is(model.Gdb.DB.First(&model.ExaSimpleUploader{}, "identifier = ? AND is_done = ?", md5, true).Error, gorm.ErrRecordNotFound)
	return err, uploads, !isDone
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: MergeFileMd5
//@description: 合并文件
//@param: md5 string, fileName string
//@return: err error

func (s *SimpleUploader) MergeFileMd5(md5 string, fileName string) (err error) {

	dir := "./chunk/" + md5
	postmap := map[string]string{}
	// 如果文件上传成功 不做后续操作 通知成功即可
	if !errors.Is(model.Gdb.DB.First(&model.ExaSimpleUploader{}, "identifier = ? AND is_done = ?", md5, true).Error, gorm.ErrRecordNotFound) {
		return nil
	}

	//获取文件类型
	tmpArr := strings.Split(fileName, ".")
	tmpLen := len(tmpArr)
	if tmpLen < 1 {
		utils.OutLog("MergeFileMd5", "文件名异常")
		return errors.New("文件名异常")
	}
	//拿后缀名
	fileType := tmpArr[tmpLen-1]

	//写入数据库名称
	md5Str := fmt.Sprintf("%d%s", time.Now().Unix(), utils.Random("all", 15))
	file_name := utils.HexMd5(md5Str) //文件名
	finishDir := utils.AppConf.Path + "/" + file_name + "/"

	pathArr := strings.Split(utils.AppConf.Path, "/")
	if len(pathArr) > 0 {
		postmap["video_url"] = "/" + pathArr[len(pathArr)-1] + "/" + file_name + "/" + file_name + "." + fileType
	}

	//tmp := strings.Replace(fileName, "."+fileType, "", -1)

	postmap["title"] = file_name

	// 打开切片文件夹
	rd, err := ioutil.ReadDir(dir)
	_ = os.MkdirAll(finishDir, os.ModePerm)
	// 创建目标文件
	fd, err := os.OpenFile(finishDir+file_name+"."+fileType, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		utils.OutLog("MergeFileMd5", "创建目标文件错误")
		return
	}
	// 关闭文件
	defer fd.Close()
	// 将切片文件按照顺序写入
	for k := range rd {
		content, _ := ioutil.ReadFile(dir + "/" + fileName + strconv.Itoa(k+1))
		_, err = fd.Write(content)
		if err != nil {
			_ = os.Remove(finishDir + fileName)
		}
	}

	if err != nil {
		utils.OutLog("MergeFileMd5", err.Error())
		return err
	}
	err = model.Gdb.DB.Transaction(func(tx *gorm.DB) error {
		// 删除切片信息
		if err = tx.Delete(&model.ExaSimpleUploader{}, "identifier = ? AND is_done = ?", md5, false).Error; err != nil {
			fmt.Println(err)
			return err
		}
		data := model.ExaSimpleUploader{
			Identifier: md5,
			IsDone:     true,
			FilePath:   finishDir + fileName,
			Filename:   fileName,
			ChunkNumber: postmap["title"],
		}
		// 添加文件信息
		if err = tx.Create(&data).Error; err != nil {
			fmt.Println(err)
			utils.OutLog("MergeFileMd5", err.Error())
			return err
		}
		return nil
	})

	err = os.RemoveAll(dir) // 清除切片

	if err == nil {
		//发送到甜心
		status, msg := PostToServer(postmap)
		if status != 200 {
			utils.OutLog("MergeFileMd5", msg)
			return errors.New(msg)
		}
	}
	return err
}

func PostToServer(mdata map[string]string) (int, string) {

	data, _ := json.Marshal(mdata)
	testAES := utils.SetAES(utils.AppConf.PrivateKey, utils.AppConf.Iv, "pkcs5", 32)
	pay_data := testAES.AesEncryptString(string(data))

	api_url := utils.AppConf.Tianxin + "/api/video_add.do"
	api_method := "POST"
	http_header := map[string]string{}
	http_header["Content-type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	http_header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"

	param := map[string]string{}
	param["proId"] = utils.AppConf.Proid
	param["apiData"] = string(pay_data)

	paramform := utils.MapCreatLinkSort(param, "&", true, false)
	utils.OutLog("PostToServer", api_url+"====>"+paramform)
	api_status, api_b := utils.HttpBody(api_url, api_method, paramform, http_header)

	if api_status != 200 {
		utils.OutLog("PostToServer", "请求错误")
		return api_status, "请求错误"
	}

	rearm := map[string]string{}
	_ = json.Unmarshal(api_b, rearm)
	utils.OutLog("PostToServer", string(api_b))
	return 200, "success"
}
