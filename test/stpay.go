package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"strings"

// 	"gitlab.stagingvip.net/publicGroup/public/common"
// )

// type STPAY struct {
// 	Mer_code   string //商户账号
// 	Return_url string //同步跳转
// 	Notify_url string //异步回调
// 	Aes_key    string //AES key
// 	Key        string //商户key
// 	Pay_url    string //支付地址
// }

// //stpay
// func StPay() (int, string, string, string, string, map[string]string) {

// 	api := STPAY{
// 		Return_url: "https://www.baidu.com/",
// 		Pay_url:    "https://apiv2.stpay.cc",
// 		Mer_code:   "cRHvEvhkzWfXZ3LpDjpgogry7gyWMxx5H4dn",
// 		Key:        "BsLHBpR9wPkMH6dnQIsFJ6hq9HvM646wrlVLBKet62dNg", //w
// 		Notify_url: "https://www.baidu.com/",
// 		Aes_key:    "tizhqeeyc2aD8EUTbR9kqJVYLJJsbA6F",
// 	}

// 	p := PayData{
// 		Amount:       "300.00",
// 		Order_number: "456287419123210342365",
// 		Pay_bank:     "ALIPAY", //支付宝扫码
// 		Ip:           "127.0.0.1",
// 	}
// 	http_header = make(map[string]string)
// 	http_header["Content-type"] = "application/x-www-form-urlencoded; charset=UTF-8"
// 	http_header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"

// 	//请求参数
// 	param_form := map[string]string{}
// 	log_path := ""
// 	api_method := "POST"
// 	re_status := 100
// 	re_msg := "请求错误"
// 	img_url := ""

// 	//请求参数
// 	action := "deposit"
// 	data_str := fmt.Sprintf(`{"RequestOrderNo":"%s","BankCode":"%s","Money":"%s","CallBackUrl":"%s","UserLevel":"1"}`, p.Order_number, p.Pay_bank, p.Amount, api.Notify_url)
// 	aes := common.SetAESECB(api.Aes_key, "", "", "hex", 32)
// 	//data值
// 	data := strings.ToUpper(aes.AesEncryptString(data_str))

// 	sign_str := fmt.Sprintf("Action=%s&Data=%s&MerchantId=%s&Key=%s", action, data, api.Mer_code, api.Key)
// 	//sign值
// 	sign := common.HexMd5(sign_str)

// 	param_result := fmt.Sprintf("MerchantId=%s&Action=%s&Data=%s&Sign=%s", api.Mer_code, action, data, sign)

// 	pay_url := fmt.Sprintf("%s/action", api.Pay_url)
// 	fmt.Println(param_result)
// 	t_status, msg_b := common.HttpBody(pay_url, api_method, param_result, http_header)
// 	common.LogsWithFileName(log_path, "stpay_create_", "param_result->"+param_result+"\ndata_str->"+data_str+"\nsign_str"+sign_str+"\nmsg->"+string(msg_b))
// 	if t_status != 200 {
// 		re_msg = "支付请求错误"
// 		return re_status, re_msg, api_method, api.Pay_url, img_url, param_form
// 	}
// 	json_res := make(map[string]interface{})
// 	err := json.Unmarshal(msg_b, &json_res)
// 	if err != nil {
// 		re_msg = err.Error()
// 		return re_status, re_msg, api_method, api.Pay_url, img_url, param_form
// 	}
// 	fmt.Println(string(msg_b))
// 	//捞取结果，赋值到变量
// 	Code := fmt.Sprintf("%v", json_res["Code"])
// 	result := fmt.Sprintf("%v", json_res["Result"])
// 	re_msg = fmt.Sprintf("%v", json_res["ErrMsg"])
// 	if Code != "0" {
// 		return re_status, re_msg, api_method, api.Pay_url, img_url, param_form
// 	}
// 	url_body := aes.AesDecryptString(result)
// 	url_res := make(map[string]interface{})
// 	err = json.Unmarshal([]byte(url_body), &url_res)

// 	if err != nil {
// 		return re_status, re_msg, api_method, api.Pay_url, img_url, param_form
// 	}
// 	img_url = url_res["Url"].(string)
// 	fmt.Println(img_url)
// 	re_status = 200
// 	re_msg = "success"
// 	return re_status, re_msg, api_method, api.Pay_url, img_url, param_form
// }
