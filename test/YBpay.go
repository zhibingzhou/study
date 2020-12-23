package main

import (
	"encoding/json"
	"fmt"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

type YBPAY struct {
	Mer_code   string `json:"mer_code"`   //商户API识别码
	Notify_url string `json:"notify_url"` //异步回调
	Key        string `json:"key"`        //商户key
	Pay_url    string `json:"pay_url"`    //支付地址
}

//代付
func YBPayFor() (int, string, string) {

	pay_data := map[string]string{
		"order_number": "1154315132ccc111c4vc4", //订单号
		"amount":       "100.00",
		"bank_branch":  "华夏银行",                //所属支行
		"bank_code":    "CCB",                 //编码
		"card_number":  "6217003590007493464", //卡号
		"card_name":    "杨雷",                  //姓名
	}

	// 代付
	api := kspay{
		Pay_url:    "http://api.ybfu8.com",
		Mer_code:   "YB2008005",
		Key:        "0f485a376a334d11e837a73c7273d853", //w
		Notify_url: "http://api.sfwage.com/",
	}
   
	api_config, _ := json.Marshal(api)
	fmt.Println(string(api_config))
	
	log_path := ""
	//定义初始值
	api_status := 100
	pay_result := "error"
	api_msg := "代付失败"
	api_method := "POST"

	//请求参数
	param_form := map[string]string{
		"merchant":   api.Mer_code,
		"customno":   pay_data["order_number"],
		"bank_title": pay_data["bank_branch"],
		"money":      pay_data["amount"], //两位小数
		"code":       pay_data["bank_code"],
		"card":       pay_data["card_number"],
		"name":       pay_data["card_name"],
		"notifyurl":  fmt.Sprintf("%s/back/ybpay.do", api.Notify_url),
	}

	api_url := fmt.Sprintf("%s/cashout", api.Pay_url)

	//拼接
	result_url := common.MapCreatLink(param_form, "merchant,customno,money,code,name,card,notifyurl", "&", 0)
	result_url = result_url + fmt.Sprintf("%s", api.Key)
	fmt.Println(result_url)

	sign := common.HexMd5(result_url)
	param_form["sign"] = sign

	param_result := common.MapCreatLinkSort(param_form, "&", true, false)

	h_status, msg_b := common.HttpBody(api_url, api_method, param_result, http_header)

	common.LogsWithFileName(log_path, "ybpay_payfor_", "param_result->"+param_result+"\napi_url->"+api_url+"\nmsg->"+string(msg_b))

	if h_status != 200 {
		return api_status, api_msg, pay_result
	}
	var json_res map[string]interface{}
	err := json.Unmarshal(msg_b, &json_res)
	if err != nil {
		api_msg = "json错误"
		return api_status, api_msg, pay_result
	}

	if fmt.Sprintf("%v", json_res["errCode"]) != "0" {
		api_msg = fmt.Sprintf("%v", json_res["errMsg"])
		return api_status, api_msg, pay_result
	}

	api_status = 200
	api_msg = "success"
	pay_result = "processing"
	return api_status, api_msg, pay_result
}

func YBPayQuery() (int, string, string) {

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

	common.LogsWithFileName(log_path, "ybpay_payfor_", "param_result->"+param_result+"\napi_url->"+api_url+"\nmsg->"+string(msg_b))
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
