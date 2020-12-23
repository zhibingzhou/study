package main

import (
	"bytes"
	"crypto/aes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

var day_f string = "20060102150405"

//xpay
type XPAY struct {
	Notify_url string            `json:"notify_url"`
	Pay_url    string            `json:"pay_url"`
	Mer_code   string            `json:"mer_code"`
	Aes_key    string            `json:"aes_key"`
	Key        string            `json:"key"`
	Header     map[string]string `json:"header"`
}

var Pay_X XPAY

func init() {

}

// Key:        "zbAykPMmXgILq4mw", //w
// Aes_key:    "v7kFdHW+gnYmIRXNahGD5w==",
// ONLINE_B2B

//xpay
func xPay() (int, string, string, string, string, map[string]string) {
	// https://trx.3rdpay.net/trx/rest/trade/onlineB2B/order  B2B
	//支付
	// {"notify_url":"https://pay.sfwage.com/back/xpay.do","pay_url":"https://trx.3rdpay.net/trx/rest/transfer/order","mer_code":"MEMM000130JJ","aes_key":"/RwAFnncBdlazN+mDFBKAA==","key":"9JLgnYzYiW4xmal0","header":null}
	//{"notify_url":"http://api.sfwage.com/call/xpay.do","pay_url":"https://trx.3rdpay.net/trx/rest/trade/cardToCard/order","mer_code":"MEMM000130JJ","aes_key":"UOhoEEMt4wCRo9SBBfkcOQ==","key":"Zua6kJ8XgnIYo6aa","header":null} 转卡飞行
	api := XPAY{
		Notify_url: "http://api.sfwage.com/call/xpay.do",
		Pay_url:    "https://trx.3rdpay.net/trx/rest/trade/btPay/order",
		Mer_code:   "MEMM000130JJ",
		Key:        "XHG9B4V4MVlexvr9", //w
		Aes_key:    "1OnLq+yvqRfPnBVwXWsQYg==",
	}

	api_config, _ := json.Marshal(api)
	fmt.Println(string(api_config))

	api.Header = make(map[string]string)
	api.Header["Content-type"] = "application/json; charset=UTF-8"
	api.Header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"

	p := PayData{
		Amount:       "300.00",
		Order_number: "45634111281165",
		Pay_bank:     "BTPAY", //CARDTOCARD
		Ip:           "127.0.0.1",
	}
	aeskey, err := base64.StdEncoding.DecodeString(api.Aes_key)
	if err != nil {
		fmt.Println("解码出错")
	}
	log_path := ""
	api_method := "POST"
	re_status := 100
	re_msg := "请求错误"
	img_url := ""

	//业务参数
	yparam_form := map[string]string{
		"merchantNo":        api.Mer_code,
		"orderNo":           p.Order_number,
		"productCode":       p.Pay_bank,
		"tradeProfitType":   "ReceivableProduct", //类型支付
		"orderAmount":       p.Amount,
		"serverCallbackUrl": api.Notify_url,
		"goodsName":         "test",
		"bankCode":          "ICBC",
		"bankBusinessType":  "B2C",
		"bankCardType":      "DEBIT",
	}

	//请求参数
	param_form := map[string]string{
		"merchantNo":       api.Mer_code,
		"orderNo":          p.Order_number,
		"tradeProfitType":  "ReceivableProduct", //类型支付
		"productCode":      p.Pay_bank,
		"bankBusinessType": "B2C",
		"bankCardType":     "DEBIT",
	}

	//json
	result_url, _ := json.Marshal(yparam_form)

	fmt.Println(string(result_url))

	mer_ecb := common.SetAESECB(string(aeskey), "", "", "", 16)
	contents := mer_ecb.AesEncryptString(string(result_url))
	param_form["content"] = contents

	fmt.Println(contents)

	//拼接
	param := MapCreatLinkSort(yparam_form, ",", true, false)
	//首尾拼接key
	param = api.Key + "," + param + "," + api.Key

	fmt.Println(param)
	//sha 256 加密
	h := sha256.New()
	h.Write([]byte(param))
	bs := fmt.Sprintf("%x", h.Sum(nil))

	fmt.Println(bs)

	param_form["sign"] = string(bs)
	params, _ := json.Marshal(param_form)
	rep := string(params)

	fmt.Println(rep)

	h_status, msg_b := common.HttpBody(api.Pay_url, api_method, rep, api.Header)

	common.LogsWithFileName(log_path, "xpay_create_", "param->"+param+"\nmsg->"+string(msg_b))
	if h_status != 200 {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	var json_res map[string]interface{}

	err = json.Unmarshal(msg_b, &json_res)
	if err != nil {
		re_msg = "json错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	//接收参数
	fmt.Println(string(msg_b))

	jsonresult, _ := json_res["content"].(string)
	//jsonresult := `423V4I0Rcau/s/d5sWc98i8V6SIDu1pc2mVzSsrNUjdlv8yKfotd0em+QQmXDDZRPEUwd5mFNdSyLiyUFDjt1MUEUZXYt1TcYpaFhCTq6n0hUjfdMk4jg3MlWtL8yMAprJWRPV4P0hSV5Kqq5A1I1HF0tcVhpglJyxgg89C1i1H/SJUSyVFNDTThC1iCBmJ7q3eF8MMqw3hnjOsv6kdjtxZaf3mIUGo46x+8mwrlnT0=`
	fmt.Println(jsonresult)

	get_ecb := common.SetAESECB(string(aeskey), "", "", "", 16)
	url_body := get_ecb.AesDecryptString(jsonresult)

	fmt.Println(url_body)

	url_res := make(map[string]interface{})
	err = json.Unmarshal([]byte(url_body), &url_res)

	if fmt.Sprintf("%v", url_res["responseCode"]) != "0000" {
		re_msg = fmt.Sprintf("%v", url_res["errorMsg"])
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	img_url = fmt.Sprintf("%v", url_res["payUrl"])

	if img_url == "" {
		re_msg = "接口错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}
	fmt.Println(img_url)
	re_status = 200
	re_msg = "success"
	return re_status, re_msg, api_method, img_url, img_url, param_form
}

//代付
func PayFor() (int, string, string) {

	log_path := ""
	// 代付
	api := XPAY{
		Notify_url: "https://pay.sfwage.com/back/xpay.do",
		Pay_url:    "https://trx.3rdpay.net/trx/rest/transfer/order",
		Mer_code:   "MEMM000130JJ",
		Key:        "9JLgnYzYiW4xmal0", //w
		Aes_key:    "/RwAFnncBdlazN+mDFBKAA==",
	}

	api_config, _ := json.Marshal(api)
	fmt.Println(string(api_config))

	api.Header = make(map[string]string)
	api.Header["Content-type"] = "application/json; charset=UTF-8"
	api.Header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"

	pay_data := map[string]string{
		"order_number": "115431513211144", //订单号
		"amount":       "100.00",
		"bank_branch":  "华夏银行",                   //所属支行
		"bank_code":    "HXB",                    //编码
		"card_number":  "6230200201123900000730", //卡号
		"card_name":    "杨雷",                     //姓名
	}

	//定义初始值
	api_status := 100
	pay_result := "error"
	api_msg := "代付失败"
	api_method := "POST"

	//业务参数
	yparam_form := map[string]string{
		"merchantNo":        api.Mer_code,
		"orderNo":           pay_data["order_number"],
		"orderAmount":       pay_data["amount"],
		"bankCode":          pay_data["bank_code"],
		"accountName":       pay_data["card_name"],
		"bankBusinessType":  "PRIVATE",
		"accountNo":         pay_data["card_number"],
		"productCode":       "PAYAPI_PRIVATE", //产品编码
		"tradeProfitType":   "PayProduct",     //业务类型
		"serverCallbackUrl": api.Notify_url,
	}

	//请求参数
	param_form := map[string]string{
		"merchantNo":      api.Mer_code,
		"orderNo":         pay_data["order_number"],
		"tradeProfitType": "PayProduct", //类型代付
		"productCode":     "PAYAPI_PRIVATE",
	}

	aeskey, err := base64.StdEncoding.DecodeString(api.Aes_key)

	if err != nil {
		api_msg = "base4 解密aeskey出错"
		return api_status, api_msg, pay_result
	}
	//json
	result_url, _ := json.Marshal(yparam_form)

	mer_ecb := common.SetAESECB(string(aeskey), "", "", "", 16)
	contents := mer_ecb.AesEncryptString(string(result_url))
	param_form["content"] = contents

	//拼接
	params := MapCreatLinkSort(yparam_form, ",", true, false)
	//首尾拼接key
	params = api.Key + "," + params + "," + api.Key

	//sha 256 加密
	h := sha256.New()
	h.Write([]byte(params))
	bs := fmt.Sprintf("%x", h.Sum(nil))

	param_form["sign"] = string(bs)
	rep, _ := json.Marshal(param_form)
	param := string(rep)

	api_status, api_b := common.HttpBody(api.Pay_url, api_method, param, api.Header)

	common.LogsWithFileName(log_path, "xpay_payfor_", "param->"+param+"\nmsg->"+string(api_b))

	fmt.Println(string(api_b))
	if api_status != 200 {
		return api_status, api_msg, pay_result
	}

	var json_res map[string]interface{}

	err = json.Unmarshal(api_b, &json_res)
	if err != nil {
		api_msg = "json错误"
		return api_status, api_msg, pay_result
	}

	//接收参数

	jsonresult := `kTq3N1UIdCRrMOzmN0kd22Ch1sxmRWkZBpMpWjIyVY53+t0DDnXEutgSDovEGriEjZDs18vCE5hlUUnoMWAMEWh1tr/dddWtxRfu9vsf6igS9HAIQp+FYYAeqsGirWjO1eqmIwP9sFgBlKYr91R+ALr1gsXEQQQByS0NRHewEExKSfIi++DcqpunC0b/nszJL3gJ2VBL7gAoleg1qF/Tjx7i8mjBZNek4UkNb5CTeT9dzQ64JwdJWcOvhFpBqO8pWW3yEcShNHX1k0usANm3AMeFpXGsyGwhg8V0HmQbMV87RpwNXgtn3OQorLH7bQRO4BrkBE92rxuWiH5LWIaCe6QWsQntNyDXm6w7/ziy1Fl4qILd4Coq5i7Ng3alwmiDmRPUU16VMPudFvx/Hc95KDcZJXrWbsvOiINqSK0Kabs=`

	get_ecb := common.SetAESECB(string(aeskey), "", "", "", 16)
	url_body := get_ecb.AesDecryptString(jsonresult)

	url_res := make(map[string]interface{})
	err = json.Unmarshal([]byte(url_body), &url_res)
	fmt.Println(url_body)
	fmt.Println(fmt.Sprintf("%v", url_res["responseCode"]), url_body)
	if fmt.Sprintf("%v", url_res["responseCode"]) != "0000" {
		api_msg = fmt.Sprintf("%v", url_res["errorMsg"])
		return api_status, api_msg, pay_result
	}

	api_status = 200
	api_msg = "success"
	pay_result = "processing"
	return api_status, api_msg, pay_result
}

func PayQuery() (int, string, string) {

	log_path := ""
	// 代付查单
	api := XPAY{
		Notify_url: "http://api.sfwage.com/call/stpay.do",
		Pay_url:    "https://trx.3rdpay.net/trx/rest/transfer/query",
		Mer_code:   "MEMM000064EF",
		Key:        "IpM97FF1ZpX9m4T3", //w
		Aes_key:    "8yiVoBMFkp4Iotzog4Qkdg==",
	}

	pay_data := map[string]string{
		"order_number": "1154351321144", //订单号
		"amount":       "10.00",
		"bank_branch":  "农业银行", //所属支行
		"bank_code":    "",     //编码
		"card_number":  "",     //卡号
		"card_name":    "",     //姓名
	}

	api.Header = make(map[string]string)
	api.Header["Content-type"] = "application/json; charset=UTF-8"
	api.Header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"

	//定义初始值
	api_status := 100
	pay_result := "processing"
	api_msg := "代付失败"
	api_method := "POST"

	//业务参数
	yparam_form := map[string]string{
		"merchantNo":      api.Mer_code,
		"orderNo":         pay_data["order_number"],
		"productCode":     "PAYAPI_PRIVATE", //产品编码
		"tradeProfitType": "ONLINE",         //类型代付
	}

	//请求参数
	param_form := map[string]string{
		"merchantNo":      api.Mer_code,
		"orderNo":         pay_data["order_number"],
		"tradeProfitType": "ONLINE", //类型代付
		"productCode":     "PAYAPI_PRIVATE",
	}

	aeskey, err := base64.StdEncoding.DecodeString(api.Aes_key)

	if err != nil {
		fmt.Println("解码出错")
	}
	//json
	result_url, _ := json.Marshal(yparam_form)

	fmt.Println(string(result_url))

	//AES 加密
	content := EcbEncrypt(result_url, aeskey)

	mer_ecb := common.SetAESECB(string(aeskey), "", "", "", 16)
	aes_res := mer_ecb.AesEncryptString(string(result_url))

	fmt.Println("系统：", aes_res)

	//base64 加密
	contents := base64.StdEncoding.EncodeToString(content)
	param_form["content"] = contents

	fmt.Println("自写", contents)

	//拼接
	params := MapCreatLinkSort(yparam_form, ",", true, false)
	//首尾拼接key
	params = api.Key + "," + params + "," + api.Key

	fmt.Println(params)

	//sha 256 加密
	h := sha256.New()
	h.Write([]byte(params))
	bs := fmt.Sprintf("%x", h.Sum(nil))

	fmt.Println(bs)

	param_form["sign"] = string(bs)
	rep, _ := json.Marshal(param_form)
	param := string(rep)

	fmt.Println(param)

	api_status, api_b := common.HttpBody(api.Pay_url, api_method, param, api.Header)

	common.LogsWithFileName(log_path, "xpay_payfor_", "param->"+param+"\nmsg->"+string(api_b))
	if api_status != 200 {
		return api_status, api_msg, pay_result
	}

	var json_res map[string]interface{}

	err = json.Unmarshal(api_b, &json_res)
	if err != nil {
		api_msg = "json错误"
		return api_status, api_msg, pay_result
	}

	//接收参数
	fmt.Println(string(api_b))

	jsonresult, _ := json_res["content"].(string)

	fmt.Println(jsonresult)
	//base64 解密
	jsonre, _ := base64.StdEncoding.DecodeString(jsonresult)
	//aes解密
	url_body := EcbDecrypt(jsonre, aeskey)

	fmt.Println(string(url_body))

	url_res := make(map[string]interface{})
	err = json.Unmarshal(url_body, &url_res)

	if fmt.Sprintf("%v", url_res["responseCode"]) != "0000" {
		api_msg = fmt.Sprintf("%v", url_res["errorMsg"])
		return api_status, api_msg, pay_result
	}
	order_status := fmt.Sprintf("%v", url_res["orderStatus"])

	switch order_status {
	case "SUCCESS":
		pay_result = "success"
	case "FAILED", "BACK":
		pay_result = "fail"
	}
	fmt.Println(order_status)
	api_status = 200
	api_msg = "success"
	return api_status, api_msg, pay_result
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//aes ecb 加密
func EcbDecrypt(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return PKCS7UnPadding(decrypted)
}

func EcbEncrypt(data, key []byte) []byte {

	block, _ := aes.NewCipher(key)
	data = PKCS7Padding(data, block.BlockSize())
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}
