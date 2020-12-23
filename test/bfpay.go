package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/url"

// 	"gitlab.stagingvip.net/publicGroup/public/common"
// )

// type BFPAY struct {
// 	Mer_code   string `json:"mer_code"`   //商户API识别码
// 	Return_url string `json:"return_url"` //同步跳转
// 	Notify_url string `json:"notify_url"` //异步回调
// 	Key        string `json:"key"`        //商户key
// 	Pay_url    string `json:"pay_url"`    //支付地址
// }

// // //百富支付
// // func bfPay() (int, string, string, string, string, map[string]string) {

// // 	http_header = make(map[string]string)
// // 	http_header["Content-type"] = "application/x-www-form-urlencoded; charset=UTF-8"
// // 	http_header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"
// // 	api := BFPAY{
// // 		Return_url: "https://www.baidu.com/",
// // 		Pay_url:    "https://api.benepay999.cc",
// // 		Mer_code:   "tdpay",
// // 		Key:        "630230ea-059d-4d01-a6c9-4fbde9ae6545", //w
// // 		Notify_url: "http://api.sfwage.com/call/bfpay.do",
// // 	}

// // 	p := PayData{
// // 		Amount:       "500.00",
// // 		Order_number: "425312q11q11341312365",
// // 		Pay_bank:     "alipay_qr",
// // 		Ip:           "127.0.0.1",
// // 	}
// // 	//fmt.Sprintf(time.Now().Unix())
// // 	log_path := ""
// // 	api_method := "POST"
// // 	re_status := 100
// // 	re_msg := "请求错误"
// // 	img_url := ""
// // 	param_form := make(map[string]string)
// // 	notify_nurl := url.QueryEscape(api.Notify_url)
// // 	sign_str := fmt.Sprintf("amount=%s&cus_code=%s&cus_order_sn=%s&notify_url=%s&payment_flag=%s", p.Amount, api.Mer_code, p.Order_number, notify_nurl, p.Pay_bank)

// // 	//sign值
// // 	sign := common.HexMd5(sign_str + "&key=" + api.Key)

// // 	param_result := fmt.Sprintf("%s&sign=%s", sign_str, sign)

// // 	post_url := fmt.Sprintf("%s/api/payment/deposit", api.Pay_url)
// // 	h_status, msg_b := common.HttpBody(post_url, api_method, param_result, http_header)
// // 	common.LogsWithFileName(log_path, "bfpay_create_", "param_result->"+param_result+"\npost_url->"+post_url+"\nmsg->"+string(msg_b))
// // 	if h_status != 200 {
// // 		return re_status, re_msg, api_method, img_url, img_url, param_form
// // 	}
// // 	var json_res map[string]interface{}
// // 	err := json.Unmarshal(msg_b, &json_res)
// // 	if err != nil {
// // 		re_msg = "JSON解析失败"
// // 		return re_status, re_msg, api_method, img_url, img_url, param_form
// // 	}
// // 	fmt.Println(string(msg_b))

// // 	re_msg = fmt.Sprintf("%v", json_res["message"])
// // 	//捞取结果，赋值到变量
// // 	result := fmt.Sprintf("%v", json_res["result"])
// // 	if result != "success" {
// // 		return re_status, re_msg, api_method, img_url, img_url, param_form

// // 	}
// // 	order_info, ok := json_res["order_info"].(map[string]interface{})
// // 	if !ok {
// // 		return re_status, re_msg, api_method, img_url, img_url, param_form
// // 	}
// // 	re_status = 200
// // 	re_msg = "success"
// // 	img_url = fmt.Sprintf("%v", order_info["payment_uri"])
// // 	fmt.Println(img_url)
// // 	return re_status, re_msg, api_method, api.Pay_url, img_url, param_form
// // }

// // var s_format string = "2006-01-02 15:04:05"

// // var log_path string
// // var pay_data map[string]string

// func init() {
// 	// 代付
// 	api = BFPAY{
// 		Return_url: "http://api.sfwage.com/public/success.do",
// 		Pay_url:    "https://api.benepay999.cc",
// 		Mer_code:   "tdpay",
// 		Key:        "630230ea-059d-4d01-a6c9-4fbde9ae6545", //w
// 		Notify_url: "https://pay.sfwage.com/back/bfpay.do",
// 	}
// 	configbf, _ := json.Marshal(api)
// 	fmt.Println(string(configbf))
// 	pay_data = map[string]string{
// 		"order_number": "115435132551144", //订单号
// 		"amount":       "10.00",
// 		"branch":       "江苏省",                //所属支行
// 		"bank_code":    "BOSHCNSH",           //编码
// 		"card_number":  "622812123112312378", //卡号
// 		"card_name":    "王锦杨",                //姓名
// 	}
// 	http_header = make(map[string]string)
// 	http_header["Content-type"] = "application/x-www-form-urlencoded; charset=UTF-8"
// 	http_header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"
// }

// //代付
// func PayFor() (int, string, string) {
// 	log_path := ""
// 	//定义初始值
// 	api_status := 100
// 	pay_result := "error"
// 	api_msg := "代付失败"

// 	//请求参数
// 	param_form := map[string]string{
// 		"cus_code":     api.Mer_code,
// 		"cus_order_sn": pay_data["order_number"],
// 		"payment_flag": "pay_webbk",
// 		"amount":       pay_data["amount"],
// 		"bank_code":    pay_data["bank_code"],
// 		"bank_account": pay_data["card_number"],
// 		"account_name": pay_data["card_name"],
// 		"notify_url":   api.Notify_url,
// 	}
// 	encode_form := map[string]string{}
// 	for key, value := range param_form {
// 		encode_form[key] = url.QueryEscape(value)
// 	}
// 	//拼接
// 	sign_str := MapCreatLinkSort(encode_form, "&", true, false)
// 	sign_str += fmt.Sprintf("&key=%s", api.Key)
// 	fmt.Println(sign_str)

// 	sign := common.HexMd5(sign_str)
// 	param_form["sign"] = sign

// 	param_result := MapCreatLinkSort(param_form, "&", true, false)

// 	api_url := fmt.Sprintf("%s/api/charge/receive", api.Pay_url)

// 	msg_b, err := HttpPostForm(api_url, param_form)
// 	common.LogsWithFileName(log_path, "bfpay_payfor_", "param_result->"+param_result+"\napi_url->"+api_url+"\nmsg->"+string(msg_b))
// 	if err != nil {
// 		return api_status, api_msg, pay_result
// 	}
// 	var json_res map[string]interface{}
// 	err = json.Unmarshal([]byte(msg_b), &json_res)
// 	if err != nil {
// 		api_msg = "json错误"
// 		return api_status, api_msg, pay_result
// 	}

// 	if fmt.Sprintf("%v", json_res["result"]) != "success" && fmt.Sprintf("%v", json_res["status"]) != "200" {
// 		api_msg = fmt.Sprintf("%v", json_res["request_data"])
// 		return api_status, api_msg, pay_result
// 	}

// 	api_status = 200
// 	api_msg = "success"
// 	pay_result = "processing"
// 	return api_status, api_msg, pay_result
// }

// func PayQuery() (int, string, string) {
// 	log_path := ""
// 	//定义初始值
// 	api_status := 100
// 	pay_result := "processing"
// 	api_msg := "代付失败"

// 	//请求参数
// 	param_form := map[string]string{
// 		"cus_code": api.Mer_code,
// 		"order_sn": pay_data["order_number"], //这个要用第三方的订单号去查
// 	}

// 	//拼接
// 	sign_str := MapCreatLinkSort(param_form, "&", true, false)
// 	sign_str += fmt.Sprintf("&key=%s", api.Key)
// 	sign := common.HexMd5(sign_str)
// 	param_form["sign"] = sign

// 	//写log
// 	param_result := MapCreatLinkSort(param_form, "&", true, false)

// 	api_url := fmt.Sprintf("%s/api/charge/info", api.Pay_url)

// 	msg_b, err := HttpPostForm(api_url, param_form)
// 	common.LogsWithFileName(log_path, "bfpay_payquery_", "param_result->"+param_result+"\napi_url->"+api_url+"\nmsg->"+string(msg_b))
// 	if err != nil {
// 		return api_status, api_msg, pay_result
// 	}

// 	var json_res map[string]interface{}
// 	err = json.Unmarshal([]byte(msg_b), &json_res)
// 	if err != nil {
// 		api_msg = "json错误"
// 		return api_status, api_msg, pay_result
// 	}

// 	fmt.Println(string(msg_b))

// 	if fmt.Sprintf("%v", json_res["result"]) != "success" && fmt.Sprintf("%v", json_res["status"]) != "200" {
// 		api_msg = fmt.Sprintf("%v", json_res["message"])
// 		return api_status, api_msg, pay_result
// 	}
// 	order_status := map[string]interface{}{}
// 	order_status = json_res["order_info"].(map[string]interface{})
// 	fmt.Println(fmt.Sprintf("%v", order_status["order_status"]))
// 	switch fmt.Sprintf("%v", order_status["order_status"]) {
// 	case "success":
// 		pay_result = "success"
// 		api_msg = "success"
// 	case "fail", "cancel":
// 		pay_result = "fail"
// 	}
// 	api_status = 200
// 	return api_status, api_msg, pay_result
// }

// // // //bfpay代付回调
// // // func Backbfpay(c *gin.Context) {
// // // 	json_res := make(map[string]interface{})
// // // 	sign := ""
// // // 	sign_str := ""
// // // 	res := "fail"
// // // 	bodystr := ""
// // // 	if c.Request.Body != nil {
// // // 		body := make([]byte, c.Request.ContentLength)
// // // 		body, err := ioutil.ReadAll(c.Request.Body)
// // // 		bodystr = string(body)
// // // 		err = json.Unmarshal(body, &json_res)
// // // 		if err == nil {

// // // 			encode_form := map[string]string{}
// // // 			for key, value := range json_res {
// // // 				encode_form[key] = url.QueryEscape(fmt.Sprintf("%v", value))
// // // 			}
// // // 			//拼接
// // // 			sign_str = MapCreatLinkSort(encode_form, "&", true, false)
// // // 			status := fmt.Sprintf("%v", json_res["tradestatus"])
// // // 			cash_status := "1"
// // // 			if status == "success" {
// // // 				cash_status = "3"
// // // 			} else if status == "failed" {
// // // 				cash_status = "9"
// // // 			}
// // // 			if cash_status != "1" {
// // // 				amount := fmt.Sprintf("%v", json_res["order_amount"])
// // // 				trade_no := fmt.Sprintf("%v", json_res["cus_order_sn"])
// // // 				pay_order := fmt.Sprintf("%v", json_res["order_sn"])
// // // 				sign = fmt.Sprintf("%v", json_res["sign"])
// // // 				is_cent := 0
// // // 				c_status, c_msg := thread.ThreadUpdateOrder(trade_no, pay_order, amount, cash_status, sign, sign_str, is_cent)
// // // 				res = c_msg
// // // 				if c_status == 200 {
// // // 					res = "SUCCESS"
// // // 				}
// // // 			}
// // // 		}
// // // 	}

// // // 	common.LogsWithFileName(log_path, "bfpay_back_", "sign->"+sign+"\nsign_str->"+sign_str+"body"+bodystr+"\nres->"+res)

// // // 	c.Writer.WriteString(res)
// // // }
