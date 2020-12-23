package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

//THPAY支付
type THPAY struct {
	Fxnotifyurl string `json:"fxnotifyurl"`
	Pay_url     string `json:"pay_url"`
	Mer_code    string `json:"mer_code"`
	Key         string `json:"key"`
}

//天地支付
func thpay() {

	tdpay := THPAY{
		Fxnotifyurl: "http://api.sfwage.com/call/thpay.do",
		Pay_url:     "https://api.thgroup888.com",
		Mer_code:    "880161",
		Key:         "0e087e230f8e05cb59f45abb101a5274",
	}
	pay_config, _ := json.Marshal(tdpay)

	fmt.Println(string(pay_config))

	p := PayData{
		Amount:       "300.00",
		Order_number: "4563412112321232342365",
		Pay_bank:     "bank_saoma",
		Ip:           "127.0.0.1",
	}

	param_form := map[string]string{
		"fxamount":    p.Amount,
		"fxorderid":   p.Order_number,
		"fxnotifyurl": tdpay.Fxnotifyurl,
		"fxid":        tdpay.Mer_code,
		"fxretype":    "0",
		"fxip":        p.Ip,
		"fxpaytype":   p.Pay_bank,
	}

	//拼接
	result_url := fmt.Sprintf("%s%s%s%s%s", param_form["fxid"], param_form["fxorderid"], param_form["fxamount"], param_form["fxnotifyurl"], tdpay.Key)

	sign := common.HexMd5(result_url)
	param_form["fxsign"] = sign

	sbForm := "<html><head></head><body onload=\"document.forms[0].submit()\"><form name=\"order\" method=\"post\" action=\"PayBuildDomain\">"
	sbForm = strings.Replace(sbForm, "PayBuildDomain", tdpay.Pay_url, -1)
	for item := range param_form {
		sbForm += "<input name=\"" + item + "\" type=\"hidden\" value=\"" + param_form[item] + "\" />"
	}
	sbForm += "</form></body></html>"

	fmt.Println(sbForm)

	// //请求三方接口
	// param := MapCreatLinkSort(param_form, "&", true)
	// h_status, msg_b := common.HttpBody(api.Pay_url, api_method, param, http_header)
	// common.LogsWithFileName(log_path, "yfpay_create_", "param->"+param+"\nmsg->"+string(msg_b))
	// if h_status != 200 {
	// 	return re_status, re_msg, api_method, img_url, img_url, param_form
	// }

	// var json_res map[string]interface{}
	// err := json.Unmarshal(msg_b, &json_res)
	// if err != nil {
	// 	re_msg = "json错误"
	// 	return re_status, re_msg, api_method, img_url, img_url, param_form
	// }

	// if fmt.Sprintf("%v", json_res["resultCode"]) != "0000" {
	// 	re_msg = fmt.Sprintf("%v", json_res["errMsg"])
	// 	return re_status, re_msg, api_method, img_url, img_url, param_form
	// }

	// img_url = fmt.Sprintf("%v", json_res["payMessage"])

	// if img_url == "" {
	// 	re_msg = "接口错误"
	// 	return re_status, re_msg, api_method, img_url, img_url, param_form
	// }

}
