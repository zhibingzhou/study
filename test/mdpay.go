package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

//马丁支付
type MDPAY struct {
	Notify_url string `json:"notify_url"`
	Pay_url    string `json:"pay_url"`
	Mer_code   string `json:"mer_code"`
	Key        string `json:"key"`
}

//马丁支付
func MdPay() (int, string, string, string, string, map[string]string) {

	log_path := ""
	api_method := "POST"
	re_status := 100
	re_msg := "请求错误"
	img_url := ""

	api := MDPAY{
		Notify_url: "http://api.sfwage.com/call/mdpay.do",
		Pay_url:    "http://51xst.cn",
		Mer_code:   "2465",
		Key:        "VfAvorTuqIz9dgxQUpKm8VG2xTQiFity",
	}

	api_config, _ := json.Marshal(api)
	fmt.Println(string(api_config))

	p := PayData{
		Amount:       "100.00",
		Order_number: "456287234495823112365",
		Pay_bank:     "PERSONAL_RED_PACK",
		Ip:           "127.0.0.1",
	}

	param_form := map[string]string{
		"mch_id":       api.Mer_code,
		"child_type":   "H5",
		"out_trade_no": p.Order_number,
		"notify_url":   url.QueryEscape(api.Notify_url),
		"pay_type":     p.Pay_bank,
		"total_fee":    p.Amount,
		"mch_secret":   api.Key,
		"timestamp":    fmt.Sprintf("%d", time.Now().Unix()),
	}

	//拼接
	result_url := common.MapCreatLinkSort(param_form, "&", true, false)
	delete(param_form, "mch_secret")

	sign := common.HexMd5(result_url)
	param_form["sign"] = strings.ToUpper(sign)

	//请求三方接口
	param := common.MapCreatLinkSort(param_form, "&", true, false)

	pay_url := fmt.Sprintf("%s/api/gateway/create", api.Pay_url)
	h_status, msg_b := common.HttpBody(pay_url, api_method, param, http_header)
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
