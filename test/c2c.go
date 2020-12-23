package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/url"

// 	"gitlab.stagingvip.net/publicGroup/public/common"
// )

// //mbpay
// type MBPAY struct {
// 	Notify_url string            `json:"notify_url"`
// 	Pay_url    string            `json:"pay_url"`
// 	SN         string            `json:"sn"`
// 	Key        string            `json:"key"`
// 	Secret_key string            `json:"secret_key"`
// 	Header     map[string]string `json:"header"`
// }

// var s_format string = "2006-01-02 15:04:05"

// var api MBPAY

// var pay_data map[string]string

// func init() {
// 	// 代付
// 	api = MBPAY{
// 		Notify_url: "https://pay.sfwage.com/back/mbpay.do",
// 		Pay_url:    "https://iccmbvip.cc/api/1",
// 		SN:         "224",
// 		Key:        "52372528de457e7d384ad8a745fe519373c991569c66b0ac036fb6c8075fbd1f", //w
// 		Secret_key: "QFXpEGFuEmPVBB90zP3OdmiX2XTcbwjf",
// 	}

// 	api_config, _ := json.Marshal(api)
// 	fmt.Println(string(api_config))

// 	pay_data = map[string]string{
// 		"order_number": "15922992898aholxkmmy", //订单号
// 		"amount":       "10.00",
// 		"branch":       "江苏省",                 //所属支行
// 		"bank_code":    "ABC",                 //编码
// 		"card_number":  "6217856000014395569", //卡号
// 		"card_name":    "小明",                  //姓名
// 		"bank_title":   "中国银行",
// 	}

// 	api.Header = make(map[string]string)
// 	api.Header["Content-type"] = "application/x-www-form-urlencoded; charset=UTF-8"
// 	api.Header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"
// }

// //代付
// func PayFor() (int, string, string) {

// 	log_path := ""
// 	//定义初始值
// 	api_status := 100
// 	pay_result := "error"
// 	api_msg := "代付失败"
// 	api_method := "POST"

// 	//请求参数
// 	param_form := map[string]string{
// 		"TransactionCode":   pay_data["order_number"],
// 		"AccountNumber":     pay_data["card_number"],
// 		"AccountName":       pay_data["card_name"],
// 		"TransactionAmount": pay_data["amount"],
// 		"BankName":          pay_data["bank_title"],
// 		"Callback":          api.Notify_url,
// 		"Version":           "0",
// 	}

// 	api.Header["Authorization"] = api.Key
// 	pay_url := api.Pay_url + fmt.Sprintf("/order/withdraw?sn=%s", api.SN)
// 	param := MapCreatLinkSort(param_form, "&", true, false)
// 	fmt.Println(param)
// 	api_status, api_b := common.HttpBody(pay_url, api_method, param, api.Header)
// 	common.LogsWithFileName(log_path, "mbpay_payfor_", "param->"+param+"\nmsg->"+string(api_b))
// 	fmt.Println(string(api_b))
// 	if api_status != 200 {
// 		return api_status, api_msg, pay_result
// 	}

// 	var json_res map[string]interface{}

// 	err := json.Unmarshal(api_b, &json_res)
// 	if err != nil {
// 		api_msg = "json错误"
// 		return api_status, api_msg, pay_result
// 	}

// 	if fmt.Sprintf("%v", json_res["status"]) != "true" {
// 		api_msg = fmt.Sprintf("%v", "下单失败")
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
// 	pay_result := "error"
// 	api_msg := "代付失败"
// 	api_method := "POST"
// 	//请求参数
// 	param_form := map[string]string{
// 		"OrderType ":      "withdraw",
// 		"TransactionCode": pay_data["order_number"],
// 		"SerialNumber":    pay_data["order_number"],
// 	}

// 	api.Header["Authorization"] = api.Key
// 	pay_url := api.Pay_url + fmt.Sprintf("/order/query?sn=%s", api.SN)
// 	param := MapCreatLinkSort(param_form, "&", true, false)
// 	fmt.Println(param)
// 	api_status, api_b := common.HttpBody(pay_url, api_method, param, api.Header)
// 	common.LogsWithFileName(log_path, "mbpay_payfor_", "param->"+param+"\nmsg->"+string(api_b))
// 	fmt.Println(string(api_b))
// 	if api_status != 200 {
// 		return api_status, api_msg, pay_result
// 	}

// 	var json_res map[string]interface{}

// 	err := json.Unmarshal(api_b, &json_res)
// 	if err != nil {
// 		api_msg = "json错误"
// 		return api_status, api_msg, pay_result
// 	}

// 	if fmt.Sprintf("%v", json_res["status"]) != "true" {
// 		api_msg = fmt.Sprintf("%v", "下单失败")
// 		return api_status, api_msg, pay_result
// 	}

// 	url_res := make(map[string]interface{})
// 	err = json.Unmarshal([]byte(fmt.Sprintf("%v", json_res["data"])), &url_res)

// 	order_status := fmt.Sprintf("%v", url_res["Status"])

// 	fmt.Println(order_status)
// 	switch order_status {
// 	case "3":
// 		pay_result = "success"
// 	case "4":
// 		pay_result = "fail"
// 	}

// 	api_status = 200
// 	api_msg = "success"
// 	pay_result = "processing"
// 	return api_status, api_msg, pay_result
// }

// //Post from 表单提交
// func HttpPostForm(post_url string, param_form map[string]string) ([]byte, error) {

// 	data := make(url.Values)
// 	for key, value := range param_form {
// 		data[key] = []string{value}
// 	}
// 	http.Header.Add(http.Header{}, "Authorization", "")
// 	//把post form 表单提交 发送给目标服务器
// 	resp, err := http.PostForm(post_url, data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return body, nil
// }

// // import (
// // 	"encoding/json"
// // 	"fmt"
// // 	"io/ioutil"
// // 	"net/http"
// // 	"net/url"
// // 	"time"

// // 	"gitlab.stagingvip.net/publicGroup/public/common"
// // )

// // //c2c
// // type C2CPAY struct {
// // 	Notify_url string
// // 	Pay_url    string
// // 	Mer_code   string
// // 	Key        string
// // }

// // var s_format string = "2006-01-02 15:04:05"

// // var api C2CPAY

// // var pay_data map[string]string

// // func init() {
// // 	// 代付
// // 	api = C2CPAY{
// // 		Notify_url: "https://www.baidu.com/",
// // 		Pay_url:    "",
// // 		Mer_code:   "27143569",
// // 		Key:        "oUSGwsy2KzNZ8MVL", //w
// // 	}

// // 	pay_data = map[string]string{
// // 		"order_number": "115435132551144", //订单号
// // 		"amount":       "10.00",
// // 		"branch":       "江苏省",                 //所属支行
// // 		"bank_code":    "ABC",                 //编码
// // 		"card_number":  "6228481199052946778", //卡号
// // 		"card_name":    "王锦杨",                 //姓名
// // 	}
// // 	http_header = make(map[string]string)
// // 	http_header["Content-type"] = "application/x-www-form-urlencoded; charset=UTF-8"
// // 	http_header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"
// // }

// // //代付
// // func PayFor() (int, string, string) {

// // 	log_path := ""

// // 	//代付请求地址
// // 	api.Pay_url = "https://api.duoduopay.cc/api/generateorder"
// // 	//定义初始值
// // 	api_status := 100
// // 	pay_result := "error"
// // 	api_msg := "代付失败"

// // 	//请求参数
// // 	param_form := map[string]string{
// // 		"merchant":        api.Mer_code,
// // 		"tradeno":         pay_data["order_number"],
// // 		"tradedate":       fmt.Sprintf(time.Now().Format(s_format)), //时间
// // 		"bankcode":        pay_data["bank_code"],
// // 		"tradedesc":       api.Mer_code,
// // 		"bankaccountno":   pay_data["card_number"],
// // 		"bankaccountname": pay_data["card_name"],
// // 		"bankaddress":     pay_data["bank_branch"],
// // 		"currency":        "CNY",
// // 		"totalamount":     pay_data["amount"],
// // 		"notifyurl":       api.Notify_url,
// // 	}

// // 	//拼接
// // 	sign_str := MapCreatLinkSort(param_form, "&", true, true)
// // 	sign_str += fmt.Sprintf("&paykey=%s", api.Key)

// // 	sign := common.HexMd5(sign_str)
// // 	param_form["sign"] = sign

// // 	param := MapCreatLinkSort(param_form, "&", true, false)

// // 	result, err := HttpPostForm(api.Pay_url, param_form)
// // 	if err != nil {
// // 		return api_status, api_msg, pay_result
// // 	}
// // 	//post form 表单提交返回值 result 里面
// // 	fmt.Println(string(result))
// // 	common.LogsWithFileName(log_path, "c2cpay_payfor_", "param->"+param+"\nmsg->"+string(result))
// // 	if err != nil {
// // 		return api_status, api_msg, pay_result
// // 	}

// // 	var json_res map[string]interface{}
// // 	err = json.Unmarshal(result, &json_res)
// // 	if err != nil {
// // 		api_msg = "json错误"
// // 		return api_status, api_msg, pay_result
// // 	}

// // 	if fmt.Sprintf("%v", json_res["resultCode"]) != "0000" && fmt.Sprintf("%v", json_res["tradestatus"]) != "SUCCESS" {
// // 		api_msg = fmt.Sprintf("%v", json_res["errMsg"])
// // 		return api_status, api_msg, pay_result
// // 	}

// // 	api_status = 200
// // 	api_msg = "success"
// // 	pay_result = "processing"
// // 	return api_status, api_msg, pay_result
// // }

// // func PayQuery() (int, string, string) {

// // 	log_path := ""

// // 	//代付查询地址
// // 	api.Pay_url = "https://api.duoduopay.cc/api/queryorder"

// // 	//定义初始值
// // 	api_status := 100
// // 	pay_result := "processing"
// // 	api_msg := "代付失败"

// // 	//请求参数
// // 	param_form := map[string]string{
// // 		"merchant": api.Mer_code,
// // 		"tradeno":  pay_data["order_number"],
// // 	}

// // 	//拼接
// // 	sign_str := MapCreatLinkSort(param_form, "&", true, true)
// // 	sign_str += fmt.Sprintf("&paykey=%s", api.Key)

// // 	sign := common.HexMd5(sign_str)
// // 	param_form["sign"] = sign

// // 	param := MapCreatLinkSort(param_form, "&", true, true)
// // 	//把post form 表单提交 发送给目标服务器
// // 	result, err := HttpPostForm(api.Pay_url, param_form)
// // 	if err != nil {
// // 		return api_status, api_msg, pay_result
// // 	}
// // 	//post form 表单提交返回值 在body 里面
// // 	fmt.Println(string(result))
// // 	common.LogsWithFileName(log_path, "c2cpay_payquery_", "param->"+param+"\nmsg->"+string(result))

// // 	var json_res map[string]interface{}
// // 	err = json.Unmarshal(result, &json_res)
// // 	if err != nil {
// // 		api_msg = "json错误"
// // 		return api_status, api_msg, pay_result
// // 	}

// // 	if fmt.Sprintf("%v", json_res["resultCode"]) != "0000" {
// // 		api_msg = fmt.Sprintf("%v", json_res["errMsg"])
// // 		return api_status, api_msg, pay_result
// // 	}

// // 	order_status := fmt.Sprintf("%v", json_res["tradestatus"])
// // 	fmt.Println(order_status)
// // 	switch order_status {
// // 	case "SUCCESS":
// // 		pay_result = "success"
// // 	case "FAILED", "NOT_EXIST":
// // 		pay_result = "fail"
// // 	}

// // 	api_status = 200
// // 	api_msg = "success"
// // 	pay_result = "processing"
// // 	return api_status, api_msg, pay_result
// // }

//Post from 表单提交
func HttpPostForm(post_url string, param_form map[string]string) ([]byte, error) {

	data := make(url.Values)
	for key, value := range param_form {
		data[key] = []string{value}
	}
	//把post form 表单提交 发送给目标服务器
	resp, err := http.PostForm(post_url, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
