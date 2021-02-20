package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

//马丁支付
type ABC struct {
	Notify_url string `json:"notify_url"`
	Pay_url    string `json:"pay_url"`
	Mer_code   string `json:"mer_code"`
	Key        string `json:"key"`
}

//马丁支付
func ABCPAY() (int, string, string, string, string, map[string]string) {

	log_path := ""
	api_method := "POST"
	re_status := 100
	re_msg := "请求错误"
	img_url := ""

	api := MDPAY{
		Notify_url: "http://api.sfwage.com/call/mdpay.do",
		Pay_url:    "https://api.quickpay123.com/api/pay/create_order",
		Mer_code:   "20000022",
		Key:        "RJTYJ3CT7K2MZ2OAIOXHUHCFH3RHIZR1SM4GO0UQJQ3OCRDJNQVALMORT3H4QVJKLDMUGCAUDPD7D6FOJODGAT3MM3KHME3T51OKZC8INYLUQOZ79PIWCJEUWSTIQTXB",
	}

	api_config, _ := json.Marshal(api)
	fmt.Println(string(api_config))

	p := PayData{
		Amount:       "50000",
		Order_number: "456287234495823112365",
		Pay_bank:     "PERSONAL_RED_PACK",
		Ip:           "127.0.0.1",
	}

	param_form := map[string]string{
		"mchId":      api.Mer_code,
		"appId":      "c61b1ac3865b4c2a8adf4c2f8beeaf74",
		"productId":  "8029",
		"mchOrderNo": p.Order_number,
		"notifyUrl":  api.Notify_url,
		"subject":    "test",
		"body":       "test",
		"currency":   "cny",
		"amount":     p.Amount,
		"reqTime":    fmt.Sprintf(time.Now().Format("20060102150405")),
		"version":    "1.0",
	}

	//拼接
	result_url := common.MapCreatLinkSort(param_form, "&", true, false)
	result_url += "&key=" + api.Key
	sign := common.HexMd5(result_url)
	param_form["sign"] = strings.ToUpper(sign)

	//请求三方接口
	param := common.MapCreatLinkSort(param_form, "&", true, false)

//	pay_url := fmt.Sprintf("%s/api/gateway/create", api.Pay_url)
	h_status, msg_b := common.HttpBody(api.Pay_url, api_method, param, http_header)
	fmt.Println(param)
	common.LogsWithFileName(log_path, "mdpay_create_", "param->"+param+"\nmsg->"+string(msg_b))
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

	if fmt.Sprintf("%v", json_res["status"]) != "true" {
		re_msg = fmt.Sprintf("%v", json_res["msg"])
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	order_info, ok := json_res["data"].(map[string]interface{})
	if !ok {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	img_url = fmt.Sprintf("%v", order_info["url"])

	if img_url == "" {
		re_msg = "接口错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}
	fmt.Println(img_url)
	re_status = 200
	re_msg = "success"
	return re_status, re_msg, api_method, img_url, img_url, param_form

}
