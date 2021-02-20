package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

//aipay
type AIPAY struct {
	Return_url string
	Pay_url    string
	Mer_code   string
	Key        string
}

//天地支付
func Aipay() (int, string, string, string, string, map[string]string) {

	api := AIPAY{
		Return_url: "http://pay.yunpays.net/call/aipay.do",
		Pay_url:    "https://api.ai-pay88.com/order_apply/",
		Mer_code:   "880113",
		Key:        "92720b2893197f8096d26f50155b4fe5",
	}
	pay_config, _ := json.Marshal(api)

	fmt.Println(string(pay_config))

	p := PayData{
		Amount:       "300.00",
		Order_number: "45634121123c21232342365",
		Pay_bank:     "zfb04",
		Ip:           "127.0.0.1",
	}

	log_path := ""
	api_method := "POST"
	re_status := 100
	re_msg := "请求错误"
	img_url := ""
	param_form := map[string]string{}
	amount, err := RmbTranfer(p.Amount, false)

	if err != nil {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}
	//请求参数
	param_form = map[string]string{
		"mchId":       api.Mer_code,
		"productId":   p.Pay_bank,
		"mchOrderNo":  p.Order_number,
		"amount":      amount,
		"clientIp":    p.Ip,
		"callbackUrl": api.Return_url,
		"reqTime":     fmt.Sprintf(time.Now().Format("20060102150405")),
	}

	rep := common.MapCreatLinkSort(param_form, "&", true, false)
	rep = rep + "&key=" + api.Key

	sign := common.HexMd5(rep)
	sign = strings.ToUpper(sign)
	param_form["sign"] = sign

	param := common.MapCreatLinkSort(param_form, "&", true, false)

	h_status, msg_b := common.HttpBody(api.Pay_url, api_method, param, http_header)
	common.LogsWithFileName(log_path, "aipay_create_", "param->"+param+"\nmsg->"+string(msg_b))
	if h_status != 200 {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	var json_res map[string]interface{}
	err = json.Unmarshal([]byte(msg_b), &json_res)
	if err != nil {
		re_msg = "json错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	if fmt.Sprintf("%v", json_res["retCode"]) != "1" {
		re_msg = fmt.Sprintf("%v", json_res["retMsg"])
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	img_url = fmt.Sprintf("%v", json_res["payUrl"])

	if img_url == "" {
		re_msg = "接口错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	re_status = 200
	re_msg = "success"
	return re_status, re_msg, api_method, img_url, img_url, param_form

}
