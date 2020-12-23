package main

import (
	"encoding/json"
	"fmt"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

//新万顺支付
type XWSPAY struct {
	Notify_url string `json:"notify_url"`
	Fxback_url string `json:"fxback_url"`
	Pay_url    string `json:"pay_url"`
	Mer_code   string `json:"mer_code"`
	Key        string `json:"key"`
}

var log_path = ""

func xwspay() (int, string, string, string, string, map[string]string) {

	api := XWSPAY{
		Notify_url: "http://api.sfwage.com/call/xwspay.do",
		Fxback_url: "http://api.sfwage.com/public/success.do",
		Pay_url:    "http://zf.hshn2020.com",
		Mer_code:   "2020111",
		Key:        "oMRcBSbJyTzwzmQBqSSnbCTAsJUvoFQh",
	}
	pay_config, _ := json.Marshal(api)

	fmt.Println(string(pay_config))

	p := PayData{
		Amount:       "300",
		Order_number: "4563412aq11239031162365",
		Pay_bank:     "alipay_wap",
		Ip:           "127.0.0.1",
	}

	api_method := "POST"
	re_status := 100
	re_msg := "请求错误"

	img_url := ""
	param_form := map[string]string{
		"fxid":        api.Mer_code,
		"fxdesc":      p.Order_number,
		"fxddh":       p.Order_number,
		"fxnotifyurl": api.Notify_url,
		"fxbackurl":   api.Fxback_url,
		"fxfee":       p.Amount,
		"fxpay":       p.Pay_bank,
		"fxip":        p.Ip,
		"fxuserid":    api.Mer_code,
	}
	img_url = fmt.Sprintf("%s/Pay", api.Pay_url)
	//拼接
	result_url := MapCreatLink(param_form, "fxid,fxdesc,fxfee,fxnotifyurl", "", 2)
	result_url += fmt.Sprintf("%s", api.Key)
	sign := common.HexMd5(result_url)
	param_form["fxsign"] = sign

	//请求三方接口
	param := MapCreatLinkSort(param_form, "&", true, false)

	//把post form 表单提交 发送给目标服务器
	result, err := HttpPostForm(img_url, param_form)
	if err != nil {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	common.LogsWithFileName(log_path, "zofpay_create_", "param->"+param+"\nmsg->"+string(result))
	if err != nil {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	var json_res map[string]interface{}
	err = json.Unmarshal(result, &json_res)
	if err != nil {
		re_msg = "json错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	if fmt.Sprintf("%v", json_res["status"]) != "1" {
		re_msg = fmt.Sprintf("%v", json_res["error"])
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	img_url = fmt.Sprintf("%v", json_res["payurl"])

	if img_url == "" {
		re_msg = "接口错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}
	fmt.Println(img_url)
	re_status = 200
	re_msg = "success"
	return re_status, re_msg, api_method, img_url, img_url, param_form
}

// /**
// * cwspay的回调接口
// * 隐式回调接口，就是服务器对服务器端的推送
//  */
// func Callcwspay(c *gin.Context) {

// 	res := "fail"
// 	sign_str := ""
// 	param := map[string]string{
// 		"callbacks":    c.PostForm("callbacks"),
// 		"appid":        c.PostForm("appid"),
// 		"pay_type":     c.PostForm("pay_type"),
// 		"success_url":  c.PostForm("success_url"),
// 		"error_url":    c.PostForm("error_url"),
// 		"out_trade_no": c.PostForm("out_trade_no"),
// 		"amount":       c.PostForm("amount"),
// 		"amount_true":  c.PostForm("amount_true"),
// 		"out_uid":      c.PostForm("out_uid"),
// 	}
// 	sign := c.PostForm("sign")
// 	if param["fxstatus"] == "succ" {
// 		sign_str := MapCreatLinkSort(param, "&", true, false)
// 		is_cent := 0
// 		call_status, _ := thread.ThreadUpdatePay(param["out_trade_no"], param["out_trade_no"], param["amount_true"], sign, sign_str, is_cent)
// 		if call_status == 200 {
// 			res = "success"
// 		}

// 	}

// 	common.LogsWithFileName(log_path, "cwspay_call_", "sign->"+sign+"\nsign_str->"+sign_str+"\nres->"+res)

// 	c.Writer.WriteString(res)
// }
