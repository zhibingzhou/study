package main

import (
	"encoding/json"
	"fmt"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

//先峰支付
type JXPAY struct {
	Pay_url    string
	Mer_code   string
	Notify_url string
	Key        string
	Return_url string
	Header     map[string]string
}

//先峰支付
func xfpay() (int, string, string, string, string, map[string]string) {
	param_form := map[string]string{}
	api := JXPAY{
		Return_url: "https://www.baidu.com/",
		Pay_url:    "http://w4.4n7s.com/",
		Mer_code:   "3073",
		Key:        "6eJvudHW", //w
		Notify_url: "https://www.baidu.com/",
	}
	api.Header = make(map[string]string)
	api.Header["Content-type"] = "application/json; charset=UTF-8"
	api.Header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"

	p := PayData{
		Amount:       "500.00",
		Order_number: "45628749123332342365",
		Pay_bank:     "alipay_static_qrcode", //支付宝扫码
		Ip:           "127.0.0.1",
	}

	log_path := ""
	api_method := "POST"
	re_status := 100
	re_msg := "请求错误"
	img_url := ""
	sign_str := fmt.Sprintf("amount=%s&merchant_user_id=%s&notify_url=%s&out_trade_no=%s&pay_way=%s&return_url=%s", p.Amount, api.Mer_code, api.Notify_url, p.Order_number, p.Pay_bank, api.Return_url)

	sign := common.HexMd5(sign_str + api.Key)

	param := fmt.Sprintf(`{"merchant_user_id":"%s","sign_type":"MD5","sign":"%s","out_trade_no":"%s","pay_way":"%s","amount":"%s","notify_url":"%s","return_url":"%s"}`, api.Mer_code, sign, p.Order_number, p.Pay_bank, p.Amount, api.Notify_url, api.Return_url)
	post_url := fmt.Sprintf("%s/gateway/merchant/order", api.Pay_url)
	h_status, msg_b := common.HttpBody(post_url, api_method, param, api.Header)

	fmt.Println(param)
	common.LogsWithFileName(log_path, "jxpay_create_", "post_url->"+post_url+"\nparam->"+param+"\nmsg->"+string(msg_b))
	if h_status != 200 {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}
	fmt.Println(string(msg_b))
	var json_res map[string]interface{}
	err := json.Unmarshal(msg_b, &json_res)
	if err != nil {
		re_msg = "json解析错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}
	code := fmt.Sprintf("%v", json_res["code"])
	re_msg = fmt.Sprintf("%v", json_res["msg"])
	if code != "200" || json_res["code"] == nil {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	data, ok := json_res["data"].(map[string]interface{})
	if !ok {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}
	re_status = 200
	re_msg = "success"
	img_url = fmt.Sprintf("%v", data["pay_url"])
	fmt.Println(img_url)
	return re_status, re_msg, api_method, img_url, img_url, param_form
}
