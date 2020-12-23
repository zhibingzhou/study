package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

type kspay struct {
	Mer_code   string `json:"mer_code"`   //商户API识别码
	Notify_url string `json:"notify_url"` //异步回调
	Key        string `json:"key"`        //商户key
	Pay_url    string `json:"pay_url"`    //支付地址
}

//卡商支付
func Kspay() (int, string, string, string, string, map[string]string) {

	http_header = make(map[string]string)
	http_header["Content-type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	http_header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"

	api := kspay{
		Pay_url:    "http://pay.shunf.net",
		Mer_code:   "tdpay",
		Key:        "PYRC7M0EukWaZBiX", //w
		Notify_url: "http://api.sfwage.com/call/kspay.do",
	}

	p := PayData{
		Amount:       "500.00",
		Order_number: "425312q11qcc113413c12365",
		Pay_bank:     "alipay_qr",
		Ip:           "127.0.0.1",
		Is_mobile:    "1",
	}

	api_config, _ := json.Marshal(api)
	fmt.Println(string(api_config))

	//fmt.Sprintf(time.Now().Unix())
	log_path := ""
	api_method := "POST"
	re_status := 100
	re_msg := "请求错误"
	img_url := ""
	param_form := make(map[string]string)
	// notify_nurl := url.QueryEscape(api.Notify_url)

	//
	rparam_form := map[string]string{
		"order_number": p.Order_number,
		"amount":       p.Amount,
		"push_url":     api.Notify_url,
		"is_show":      "",
		"is_mobile":    p.Is_mobile,
	}

	data, _ := json.Marshal(rparam_form)

	post_url := fmt.Sprintf("%s/pay/create.do", api.Pay_url)

	testAES := common.SetAES(api.Key, "", "pkcs5", 16)
	pay_data := testAES.AesEncryptString(string(data))

	param_result := fmt.Sprintf("mer_code=%s&pay_data=%s", api.Mer_code, pay_data)
	param_form["pay_data"] = pay_data
	param_form["mer_code"] = api.Mer_code

	//h_status, msg_b := common.HttpBody(post_url, api_method, param_result, http_header)

	common.LogsWithFileName(log_path, "kspay_create_", "param_result->"+param_result+"\npost_url->"+post_url)

	sbForm := "<html><head></head><body onload=\"document.forms[0].submit()\"><form name=\"order\" method=\"post\" action=\"PayBuildDomain\">"
	sbForm = strings.Replace(sbForm, "PayBuildDomain", post_url, -1)
	for item := range param_form {
		sbForm += "<input name=\"" + item + "\" type=\"hidden\" value=\"" + param_form[item] + "\" />"
	}
	sbForm += "</form></body></html>"

	fmt.Println(sbForm)

	// if h_status != 200 {
	// 	return re_status, re_msg, api_method, img_url, img_url, param_form
	// }

	// var json_res map[string]interface{}
	// err := json.Unmarshal(msg_b, &json_res)
	// if err != nil {
	// 	re_msg = "JSON解析失败"
	// 	return re_status, re_msg, api_method, img_url, img_url, param_form
	// }
	// fmt.Println(string(msg_b))

	// re_msg = fmt.Sprintf("%v", json_res["message"])
	// //捞取结果，赋值到变量
	// result := fmt.Sprintf("%v", json_res["result"])
	// if result != "success" {
	// 	return re_status, re_msg, api_method, img_url, img_url, param_form

	// }
	// order_info, ok := json_res["order_info"].(map[string]interface{})
	// if !ok {
	// 	return re_status, re_msg, api_method, img_url, img_url, param_form
	// }
	// re_status = 200
	// re_msg = "success"
	// img_url = fmt.Sprintf("%v", order_info["payment_uri"])
	// fmt.Println(img_url)
	return re_status, re_msg, api_method, api.Pay_url, img_url, param_form
}

//代付
func KsPayFor() (int, string, string) {

	pay_data := map[string]string{
		"order_number": "1154315132cc11144", //订单号
		"amount":       "500.00",
		"bank_branch":  "华夏银行",                   //所属支行
		"bank_code":    "HXB",                    //编码
		"card_number":  "6230200201123900000730", //卡号
		"card_name":    "杨雷",                     //姓名
	}

	// 代付
	api := kspay{
		Pay_url:    "http://pay.shunf.net",
		Mer_code:   "tdpay",
		Key:        "PYRC7M0EukWaZBiX", //w
		Notify_url: "http://api.sfwage.com/call/kspay.do",
	}

	log_path := ""
	//定义初始值
	api_status := 100
	pay_result := "error"
	api_msg := "代付失败"
	api_method := "POST"

	//请求参数
	param_form := map[string]string{
		"cus_code":     api.Mer_code,
		"order_number": pay_data["order_number"],
		"bank_title":   pay_data["bank_branch"],
		"amount":       pay_data["amount"],
		"bank_code":    pay_data["bank_code"],
		"card_number":  pay_data["card_number"],
		"card_name":    pay_data["card_name"],
		"notify_url":   api.Notify_url,
	}

	data, _ := json.Marshal(param_form)

	api_url := fmt.Sprintf("%s/pay/df_pay.do", api.Pay_url)

	testAES := common.SetAES(api.Key, "", "pkcs5", 16)
	post_data := testAES.AesEncryptString(string(data))

	param_result := fmt.Sprintf("mer_code=%s&pay_data=%s", api.Mer_code, post_data)

	h_status, msg_b := common.HttpBody(api_url, api_method, param_result, http_header)

	common.LogsWithFileName(log_path, "kspay_payfor_", "param_result->"+param_result+"\napi_url->"+api_url+"\nmsg->"+string(msg_b))
	if h_status != 200 {
		return api_status, api_msg, pay_result
	}
	var json_res map[string]interface{}
	err := json.Unmarshal(msg_b, &json_res)
	if err != nil {
		api_msg = "json错误"
		return api_status, api_msg, pay_result
	}

	if fmt.Sprintf("%v", json_res["Status"]) != "200" {
		api_msg = fmt.Sprintf("%v", json_res["Msg"])
		return api_status, api_msg, pay_result
	}

	api_status = 200
	api_msg = "success"
	pay_result = "processing"
	return api_status, api_msg, pay_result
}

func KSPayQuery() (int, string, string) {

	pay_data := map[string]string{
		"order_number": "115431513211144", //订单号
		"amount":       "500.00",
		"bank_branch":  "建设银行",                   //所属支行
		"bank_code":    "CCB",                    //编码
		"card_number":  "6230200201123900000730", //卡号
		"card_name":    "杨雷",                     //姓名
	}

	// 代付
	api := kspay{
		Pay_url:    "http://pay.shunf.net",
		Mer_code:   "tdpay",
		Key:        "PYRC7M0EukWaZBiX", //w
		Notify_url: "http://api.sfwage.com/call/kspay.do",
	}

	log_path := ""
	//定义初始值
	api_status := 100
	pay_result := "processing"
	api_msg := "代付失败"
	api_method := "POST"

	//请求参数
	param_form := map[string]string{
		"order_number": pay_data["order_number"],
	}

	data, _ := json.Marshal(param_form)

	api_url := fmt.Sprintf("%s/pay/cash_order.do", api.Pay_url)

	testAES := common.SetAES(api.Key, "", "pkcs5", 16)
	post_data := testAES.AesEncryptString(string(data))

	param_result := fmt.Sprintf("mer_code=%s&pay_data=%s", api.Mer_code, post_data)

	h_status, msg_b := common.HttpBody(api_url, api_method, param_result, http_header)

	common.LogsWithFileName(log_path, "kspay_payfor_", "param_result->"+param_result+"\napi_url->"+api_url+"\nmsg->"+string(msg_b))
	fmt.Println(string(msg_b))
	if h_status != 200 {
		return api_status, api_msg, pay_result
	}
	var json_res map[string]interface{}
	err := json.Unmarshal(msg_b, &json_res)
	if err != nil {
		api_msg = "json错误"
		return api_status, api_msg, pay_result
	}

	json_data, ok := json_res["Data"].(map[string]interface{})

	if !ok {
		api_msg = "json错误"
		return api_status, api_msg, pay_result
	}

	json_info, ok := json_data["order_info"].(map[string]interface{})

	if !ok {
		api_msg = "json错误"
		return api_status, api_msg, pay_result
	}

	switch fmt.Sprintf("%s", json_info["status"]) {
	case "3":
		pay_result = "success"
		api_msg = "success"
	case "9":
		pay_result = "fail"
	}

	api_status = 200
	return api_status, api_msg, pay_result
}
