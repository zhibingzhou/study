package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

//新先锋支付
type XXFPAY struct {
	Return_url string
	Pay_url    string
	Mer_code   string
	Token      string
	Notify_url string
}

//新先锋支付
func xxfPAY() (int, string, string, string, string, map[string]string) {

	http_header = make(map[string]string)
	http_header["Content-type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	http_header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"

	api := XXFPAY{
		Return_url: "https://www.baidu.com/",
		Pay_url:    "http://47.91.165.85/api/v1/charges",
		Mer_code:   "448850489841287168",
		Token:      "1171f6a833a8424bafc4bc540c7152a4", //w
		Notify_url: "https://www.baidu.com/",
	}

	p := PayData{
		Amount:       "300.00",
		Order_number: "45628749123332342365",
		Pay_bank:     "alipay_zklzz", //支付宝扫码
		Ip:           "127.0.0.1",
	}
	//fmt.Sprintf(time.Now().Unix())
	log_path := ""
	api_method := "POST"
	re_status := 100
	re_msg := "请求错误"
	img_url := ""

	param_form := map[string]string{
		"token":      api.Token,
		"money":      p.Amount,
		"outTradeNo": p.Order_number,
		"channel":    p.Pay_bank,
		"timestamp":  fmt.Sprintf("%v", time.Now().UnixNano()/1e6),
		"uid":        api.Mer_code,
		"returnUrl":  api.Return_url,
		"notifyUrl":  api.Notify_url,
	}

	//拼接
	result_url := MapCreatLinkSort(param_form, "&", true, false)
	fmt.Println(result_url)

	sign := common.HexMd5(result_url)
	sign = strings.ToUpper(sign)
	param_form["sign"] = sign
	delete(param_form, "token")
	//请求三方接口
	param := MapCreatLinkSort(param_form, "&", true, false)
	h_status, msg_b := common.HttpBody(api.Pay_url, api_method, param, http_header)
	fmt.Println(param)
	common.LogsWithFileName(log_path, "xxfpay_create_", "param->"+param+"\nmsg->"+string(msg_b))
	if h_status != 200 {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}
	fmt.Println(string(msg_b))
	var json_res map[string]interface{}
	err := json.Unmarshal(msg_b, &json_res)
	if err != nil {
		re_msg = "json错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	if fmt.Sprintf("%v", json_res["code"]) != "0" {
		re_msg = fmt.Sprintf("%v", json_res["msg"])
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	order_info, ok := json_res["data"].(map[string]interface{})
	if !ok {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	img_url = fmt.Sprintf("%v", order_info["payUrl"])

	if img_url == "" {
		re_msg = "接口错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}
	fmt.Println(img_url)
	re_status = 200
	re_msg = "success"
	return re_status, re_msg, api_method, img_url, img_url, param_form
}
