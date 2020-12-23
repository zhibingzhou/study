package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

//超凡支付
type CFPAY struct {
	Mer_code   string `json:"mer_code"`   //商户API识别码
	Return_url string `json:"return_url"` //同步跳转
	Notify_url string `json:"notify_url"` //异步回调
	Key        string `json:"key"`        //商户key
	Pay_url    string `json:"pay_url"`    //支付地址
}

//超凡支付
func cfpay() (int, string, string, string, string, map[string]string) {

	rmb, _ := RmbTranfer("500000", true)
	fmt.Println(rmb)

	http_header = make(map[string]string)
	http_header["Content-type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	http_header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"

	api := CFPAY{
		Return_url: "http://api.sfwage.com/public/success.do",
		Pay_url:    "http://47.75.96.143:9100/gateway",
		Mer_code:   "200280",
		Key:        "1a196c2c8b6a7625d24390fac143db49", //w
		Notify_url: "http://api.sfwage.com/call/cfpay.do",
	}

	api_config, _ := json.Marshal(api)
	fmt.Println(string(api_config))

	p := PayData{
		Amount:       "500.00",
		Order_number: "45628740123ca31c1c3234236c5",
		Pay_bank:     "pay.qq.wap", //支付宝扫码
		Ip:           "112.201.75.121",
	}

	//log_path := ""
	api_method := "POST"
	re_status := 100
	re_msg := "请求错误"
	img_url := ""
	log_path := ""
	param_form := map[string]string{}

	amount, err := RmbTranfer(p.Amount, false)
	if err != nil {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	param_form = map[string]string{
		"amount":     amount,
		"orderNo":    p.Order_number,
		"notifyUrl":  api.Notify_url,
		"merchantId": api.Mer_code,
		"version":    "1.0",
		"clientIp":   p.Ip,
		"service":    p.Pay_bank,
		"key":        api.Key,
		"tradeDate":  fmt.Sprintf(time.Now().Format("20060102")),
		"tradeTime":  fmt.Sprintf(time.Now().Format("150405")),
	}

	//拼接
	result_url := MapCreatLinkSort(param_form, "&", true, false)

	sign := common.HexMd5(result_url)
	fmt.Println(result_url)

	param_form["sign"] = sign
	delete(param_form, "key")

	param := MapCreatLinkSort(param_form, "&", true, false)
	h_status, msg_b := common.HttpBody(api.Pay_url, api_method, param, http_header)

	fmt.Println(api.Pay_url, param)

	common.LogsWithFileName(log_path, "cfpay_create_", "param->"+param+"\nmsg->"+string(msg_b))
	if h_status != 200 {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}
	fmt.Println(string(msg_b))

	var json_res map[string]string
	json_res = UrlToMap(string(msg_b))
	if fmt.Sprintf("%v", json_res["repCode"]) != "0001" {
		re_msg = fmt.Sprintf("%v", json_res["repMsg"])
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	img_url = fmt.Sprintf("%v", json_res["resultUrl"])

	if img_url == "" {
		re_msg = "接口错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	re_status = 200
	re_msg = "success"

	fmt.Println(img_url)
	return re_status, re_msg, api_method, img_url, img_url, param_form
}

func UrlToMap(url string) map[string]string {
	result := make(map[string]string)
	parm := strings.Split(url, "&")
	for _, value := range parm {
		keys := strings.SplitN(value, "=",2)
		result[keys[0]] = keys[1]
	}
	return result
}

//人民转换
//dic 分转元  传true ，元转分  传false
func RmbTranfer(amount string, dic bool) (string, error) {
	ramount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return "", err
	}
	if dic == false {
		ramount = ramount * 100 // 元转分
		return strconv.FormatFloat(float64(ramount), 'f', 0, 64), nil
	} else {
		ramount = ramount / 100 // 分转元,保留两们小数
		return strconv.FormatFloat(float64(ramount), 'f', 2, 64), nil
	}

}
