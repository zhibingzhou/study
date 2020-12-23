package main

import (
	"encoding/json"
	"fmt"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

//
type HYTPAY struct {
	Pay_url    string
	Mer_code   string
	Notify_url string
	Key        string
	Back_url   string
	Call_url   string
	Header     map[string]string
}

func hytpay() (int, string, string, string, string, map[string]string) {

	api_method := "POST"
	re_status := 100
	param_form := map[string]string{}
	re_msg := "请求错误"
	log_path := ""
	api := HYTPAY{
		Notify_url: "http://api.sfwage.com/call/hytpay.do",
		Back_url:   "http://api.sfwage.com/public/success.do",
		Call_url:   "http://api.sfwage.com/back/hytpay.do",
		Pay_url:    "http://service.88wu1i81r19.ofxtw.com",
		Mer_code:   "2800",
		Key:        "k5YCADxp1jfO57t9R0B883RA97mVQ0r2a",
	}
	api.Header = make(map[string]string)
	api.Header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
	api.Header["Content-Type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	p := PayData{
		Amount:       "200.00",
		Order_number: "456341332112333231362365",
		Pay_bank:     "wx2",
		Ip:           "127.0.0.1",
	}

	img_url := ""

	param := fmt.Sprintf("Amount=%s&Ip=%s&MerchantId=%s&MerchantUniqueOrderId=%s&NotifyUrl=%s&PayTypeId=%s&ReturnUrl=%s", p.Amount, p.Ip, api.Mer_code, p.Order_number, api.Notify_url, p.Pay_bank, api.Back_url)

	api_url := fmt.Sprintf("%s/InterfaceV4/CreatePayOrder/", api.Pay_url)

	api_status, api_b := common.HttpBody(api_url, api_method, param, api.Header)
	common.LogsWithFileName(log_path, "hytpay_create_", "param->"+param+"\nmsg->"+string(api_b))
	if api_status != 200 {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}
	fmt.Println(string(api_b))
	var json_res map[string]interface{}
	err := json.Unmarshal(api_b, &json_res)
	if err != nil {
		re_msg = "json错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	if fmt.Sprintf("%v", json_res["Code"]) != "0" {
		re_msg = fmt.Sprintf("%v", json_res["MessageForUser"])
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	img_url = fmt.Sprintf("%v", json_res["Url"])

	if img_url == "" {
		re_msg = "接口错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	fmt.Println(img_url)
	return re_status, re_msg, api_method, img_url, img_url, param_form
}
