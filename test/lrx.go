package main

import (
	"encoding/json"
	"fmt"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

//两人行
type LRXPAY struct {
	Notify_url string            `json:"notify_url"`
	Pay_url    string            `json:"pay_url"`
	Mer_code   string            `json:"mer_code"`
	Key        string            `json:"key"`
	Header     map[string]string `json:"header"`
}

//代付
func LRXPayFor() (int, string, string) {

	log_path := ""
	// 代付
	api := LRXPAY{
		Notify_url: "https://pay.sfwage.com/back/lrxpay.do",
		Pay_url:    "http://mch_api.paymoneyplayer.com",
		Mer_code:   "MCH-TVtCGSnfLOvoKybz",
		Key:        "oxo-slj81e7Txb0CDgHGyQLm", //w
	}

	api_config, _ := json.Marshal(api)
	fmt.Println(string(api_config))
	

	pay_data := map[string]string{
		"order_number": "1154351325c51144", //订单号
		"amount":       "10.00",
		"branch":       "江苏省",                 //所属支行
		"bank_code":    "ABC",                 //编码
		"card_number":  "6228481199052946778", //卡号
		"card_name":    "小红",
		"bank_title":   "中国建设银行", //银行名称
	}



	http_header := make(map[string]string)
	http_header["Content-type"] = "application/json; charset=UTF-8"
	http_header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"

	//定义初始值
	api_status := 100
	pay_result := "error"
	api_msg := "代付失败"
	api_method := "POST"

	amount, err := RmbTranfer(pay_data["amount"], false)
	if err != nil {
		api_msg = "金额错误"
		return api_status, api_msg, pay_result
	}
	//请求参数
	param_form := map[string]string{
		"mchId":         api.Mer_code,
		"transactionId": pay_data["order_number"],
		"bankCardNo":    pay_data["card_number"],
		"holder":        pay_data["card_name"],
		"bank":          pay_data["bank_title"],
		"bizId":         "20",
		"amount":        amount, //分
		"callbackUrl":   api.Notify_url,
	}

	//拼接
	sign_str := common.MapCreatLinkSort(param_form, "&", true, false)
	sign_str = fmt.Sprintf("%s&", api.Key) + sign_str

	sign := common.HexMd5(sign_str)
	param_form["sign"] = sign

	rep, _ := json.Marshal(param_form)
	param := string(rep)

	api_url := fmt.Sprintf("%s/api/v1/merchant/withdraws", api.Pay_url)

	api_status, api_b := common.HttpBody(api_url, api_method, param, http_header)

	common.LogsWithFileName(log_path, "lrxpay_payfor_", "param->"+param+"\nmsg->"+string(api_b)+"\napi_url->"+api_url)
	if api_status != 200 {
		return api_status, api_msg, pay_result
	}

	var json_res map[string]interface{}
	err = json.Unmarshal(api_b, &json_res)
	if err != nil {
		api_msg = "json错误"
		return api_status, api_msg, pay_result
	}

	if fmt.Sprintf("%v", json_res["status"]) != "0" && fmt.Sprintf("%v", json_res["status"]) != "2" {
		api_msg = fmt.Sprintf("%v", "下单失败")
		return api_status, api_msg, pay_result
	}

	api_status = 200
	api_msg = "success"
	pay_result = "processing"
	return api_status, api_msg, pay_result
}

func LRXPayQuery() (int, string, string) {

	log_path := ""

	// 代付
	api := LRXPAY{
		Notify_url: "https://www.baidu.com/",
		Pay_url:    "",
		Mer_code:   "27143569",
		Key:        "oUSGwsy2KzNZ8MVL", //w
	}

	pay_data := map[string]string{
		"order_number": "115435132551144", //订单号
		"amount":       "10.00",
		"branch":       "江苏省",                 //所属支行
		"bank_code":    "ABC",                 //编码
		"card_number":  "6228481199052946778", //卡号
		"card_name":    "王锦杨",                 //姓名
	}

	http_header = make(map[string]string)
	api.Header["Content-type"] = "application/json; charset=UTF-8"
	api.Header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"

	//代付查询地址
	api.Pay_url = "https://api.duoduopay.cc/api/queryorder"

	//定义初始值
	api_status := 100
	pay_result := "processing"
	api_msg := "代付失败"

	//请求参数
	param_form := map[string]string{
		"merchant": api.Mer_code,
		"tradeno":  pay_data["order_number"],
	}

	//拼接
	sign_str := MapCreatLinkSort(param_form, "&", true, true)
	sign_str += fmt.Sprintf("&paykey=%s", api.Key)

	sign := common.HexMd5(sign_str)
	param_form["sign"] = sign

	param := MapCreatLinkSort(param_form, "&", true, true)
	//把post form 表单提交 发送给目标服务器
	result, err := HttpPostForm(api.Pay_url, param_form)
	if err != nil {
		return api_status, api_msg, pay_result
	}
	//post form 表单提交返回值 在body 里面
	fmt.Println(string(result))
	common.LogsWithFileName(log_path, "c2cpay_payquery_", "param->"+param+"\nmsg->"+string(result))

	var json_res map[string]interface{}
	err = json.Unmarshal(result, &json_res)
	if err != nil {
		api_msg = "json错误"
		return api_status, api_msg, pay_result
	}

	if fmt.Sprintf("%v", json_res["resultCode"]) != "0000" {
		api_msg = fmt.Sprintf("%v", json_res["errMsg"])
		return api_status, api_msg, pay_result
	}

	order_status := fmt.Sprintf("%v", json_res["tradestatus"])
	fmt.Println(order_status)
	switch order_status {
	case "SUCCESS":
		pay_result = "success"
	case "FAILED", "NOT_EXIST":
		pay_result = "fail"
	}

	api_status = 200
	api_msg = "success"
	pay_result = "processing"
	return api_status, api_msg, pay_result
}
