package main

import (
	"encoding/json"
	"fmt"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

//STpay
type STPAY struct {
	Mer_code   string `json:"mer_code"`   //商户API识别码
	AppId      string `json:"appId"`      //商户API识别码
	Notify_url string `json:"notify_url"` //异步回调
	Key        string `json:"key"`        //商户key
	Pay_url    string `json:"pay_url"`    //支付地址
}

//三河支付
func shpay() (int, string, string, string, string, map[string]string) {

	api := STPAY{
		Pay_url:    "http://payapi.3hmo8xcf.stpay01.com:6088/",
		Mer_code:   "1112",
		Key:        "KLUTJNZL0B0GRD0NXIQYJIHJE1P7UWYTIUBPYQ1F9TZUV5LSUUKJ5JCRUSWOF0Z4TL9JSHNQ5YXA4TERT34CGHPKLDLF1JI5A3VBMFKWBHGJKGVCKMNUICCMARJZQVMA", //w
		Notify_url: "http://pay.yunpays.net/call/stpay.do",
		AppId:      "eb369eac86734b8784d7d7e28c10798b",
	}

	api_config, _ := json.Marshal(api)
	fmt.Println(string(api_config))

	p := PayData{
		Amount:       "500.00",
		Order_number: "45628740123ca31c13c32",
		Pay_bank:     "8016", //支付宝扫码
		Ip:           "127.0.0.1",
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
		"mchOrderNo": p.Order_number,
		"appId":      api.AppId,
		"mchId":      api.Mer_code,
		"clientIp":   p.Ip,
		"productId":  p.Pay_bank,
		"currency":   "cny",
		"notifyUrl":  api.Notify_url,
		"subject":    "test",
		"body":       "testbody",
	}

	//拼接
	result_url := MapCreatLinkSort(param_form, "&", true, false)
	result_url = result_url + "&key=" + api.Key

	sign := common.HexMd5(result_url)

	param_form["sign"] = sign

	url := fmt.Sprintf("%s/api/pay/create_payorder", api.Pay_url)

	param := MapCreatLinkSort(param_form, "&", true, false)

	fmt.Println(param)

	h_status, msg_b := common.HttpBody(url, api_method, param, http_header)

	common.LogsWithFileName(log_path, "stpay_create_", "url->"+url+"param->"+string(param)+"\nmsg->"+string(msg_b))
    msg_b = []byte(`{"payOrderId":"OBPAY1608367938743064","sign":"A689BBE888EF20B260E46578C7569DA5","payParams":{"payUrl":"http://syt.08iyackz.ympay.cc:8088/pays/banktransfersimple.html?amount=50000&mchId=1002&mchOrderNo=OBPAY1608367938743064&notifyUrl=http://payapi.3hmo8xcf.stpay01.com:6088/notify/obsystem/notify_res.htm&type=2&sign=44740EE478A5BC72D1B49FA40C2A34AB"},"retCode":"SUCCESS","retMsg":"下单成功"}`)
	fmt.Println(string(msg_b))
	if h_status != 200 {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	var json_res map[string]interface{}
	err = json.Unmarshal(msg_b, &json_res)
	if err != nil {
		re_msg = "json错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	if fmt.Sprintf("%v", json_res["retCode"]) != "SUCCESS" {
		re_msg = fmt.Sprintf("%v", json_res["retMsg"])
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	jsonresults := json_res["payParams"].(map[string]interface{})

	img_url = fmt.Sprintf("%v", jsonresults["payUrl"])

	if img_url == "" {
		re_msg = "接口错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	re_status = 200
	re_msg = "success"

	return re_status, re_msg, api_method, img_url, img_url, param_form
}
