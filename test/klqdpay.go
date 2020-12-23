package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

//快乐抢单
type KLQD struct {
	Pay_url  string            `json:"pay_url"`
	Mer_code string            `json:"mer_code"`
	Key      string            `json:"key"`
	Header   map[string]string `json:"header"`
}

//快乐抢单
func klqdpay() (int, string, string, string, string, map[string]string) {

	var s_format string = "2006-01-02 15:04:05"

	log_path := ""
	api_method := "POST"
	re_status := 100
	re_msg := "请求错误"
	img_url := ""

	api := KLQD{
		Pay_url:  "http://order.klqd.vip",
		Mer_code: "112",
		Key:      "e0ca2eaf8c9ce09cc55741ad94e9b561",
	}

	api.Header = make(map[string]string)
	api.Header["Content-type"] = "application/json; charset=UTF-8"
	api.Header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"

	api_config, _ := json.Marshal(api)
	fmt.Println(string(api_config))

	p := PayData{
		Amount:       "100",
		Order_number: "45628723c44c95c823c112365",
		Pay_bank:     "3",
		Ip:           "127.0.0.1",
	}

	param_form := map[string]string{
		"myId":      api.Mer_code,
		"userId":    p.Order_number,
		"orderNo":   p.Order_number,
		"payType":   p.Pay_bank,
		"amount":    p.Amount,
		"ordertime": fmt.Sprintf(time.Now().Format(s_format)), //时间
	}

	//拼接
	result_url := fmt.Sprintf("amount=%s&orderNo=%s&payType=%s&ordertime=%s&userId=%s&myId=%s",
		param_form["amount"], param_form["orderNo"], param_form["payType"], param_form["ordertime"], param_form["userId"], param_form["myId"])
	result_url += "&key=" + api.Key
	fmt.Println(result_url)
	sign := common.HexMd5(result_url)

	param_form["sign"] = strings.ToLower(sign)

	params, _ := json.Marshal(param_form)
	rep := string(params)

	//请求三方接口
	param := common.MapCreatLinkSort(param_form, "&", true, false)

	pay_url := fmt.Sprintf("%s/api.php", api.Pay_url)
	h_status, msg_b := common.HttpBody(pay_url, api_method, rep, api.Header)
	if h_status != 200 {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}
	common.LogsWithFileName(log_path, "klqdpay_create_", "param->"+param+"\nmsg->"+string(msg_b))
	result := string(msg_b)
	if result == "erro" {
		re_msg = fmt.Sprintf("%v", result)
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	img_url = fmt.Sprintf("%v", result)

	if img_url == "" {
		re_msg = "接口错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}
	fmt.Println(img_url)
	re_status = 200
	re_msg = "success"
	return re_status, re_msg, api_method, img_url, img_url, param_form

}
