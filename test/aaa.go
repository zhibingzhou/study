package main

import (
	"encoding/json"
	"fmt"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

func GetShorUrl(long_url, key string) string {

	url := "http://api.985.so/api.php?format=json"
	api_method := "GET"
	var http_header map[string]string
	http_header = make(map[string]string)
	http_header["Content-type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	http_header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"

	url = url + fmt.Sprintf("&url=%s&apikey=%s", long_url, key)

	fmt.Println(url)
	t_status, msg_b := common.HttpBody(url, api_method, "", http_header)
	common.LogsWithFileName(log_path, "GetShorUrl", "请求短链接失败！"+string(msg_b))
	if t_status != 200 {
		return ""
	}

	var json_res map[string]interface{}
	err := json.Unmarshal([]byte(msg_b), &json_res)
	if err != nil {
		return ""
	}

	if fmt.Sprintf("%v", json_res["status"]) != "1" {
		common.LogsWithFileName(log_path, "GetShorUrl", "请求短链接失败！"+fmt.Sprintf("%v", json_res["err"]))
		return ""
	}

	return fmt.Sprintf("%v", json_res["url"])
}
