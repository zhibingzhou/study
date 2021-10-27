package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDir
//@description: 批量创建文件夹
//@param: dirs ...string
//@return: err error

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			//global.GVA_LOG.Debug("create directory" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				//global.GVA_LOG.Error("create directory"+v, zap.Any(" error:", err))
				return err
			}
		}
	}
	return err
}

//创建目录
func Create(_dir string) bool {
	result := false
	exist, err := PathExists(_dir)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return result
	}

	if exist {
		fmt.Printf("has dir![%v]\n", _dir)
		return true
	} else {
		fmt.Printf("no dir![%v]\n", _dir)
		// 创建文件夹
		err := os.Mkdir(_dir, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
			return true
		}
	}
	return result
}

//判断类型
func CheckPath(header *multipart.FileHeader) (bool, string) {
	filterArr := []string{"mp4", "flv", "avi", "rmvb", "wmv"}
	tmpFile := header.Filename
	msg := "success"
	//获取文件类型
	tmpArr := strings.Split(tmpFile, ".")
	tmpLen := len(tmpArr)
	if tmpLen < 1 {
		msg = "文件名异常"
		return false, msg
	}
	fileType := tmpArr[tmpLen-1]
	if !Arr_In(filterArr, fileType) {
		msg = "文件类型不合法"
		return false, msg
	}
	return true, msg
}

/**
* 是否包含在数组中
 */
func Arr_In(arr []string, sep_str string) bool {
	res := false
	if len(arr) < 1 {
		return res
	}
	for _, arr_str := range arr {
		if arr_str == sep_str {
			res = true
			break
		}
	}
	return res
}

func OutLog(logName, logStr string) {
	LogsWithFileName(AppConf.LogPath, logName, logStr)
}

/**
* 指定文件名写入文件
 */
func LogsWithFileName(log_path, FiLeName string, msg string) {
	if IsDirExists(log_path) != true {
		os.MkdirAll(log_path, 0777)
	}
	isEnd := strings.HasSuffix(log_path, "/")
	if isEnd != true {
		log_path = log_path + "/"
	}
	timeStr := time.Now().Format("2006-01-02")
	logFile := log_path + FiLeName + timeStr + ".log"

	fout, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println(log_path, err)
		return
	}

	fout.WriteString(time.Now().Format("2006-01-02 15:04:05") + "\r\n" + msg + "\r\n=====================\r\n")
	defer fout.Close()
}

/**
* 判断目录是否存在
 */
func IsDirExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}
