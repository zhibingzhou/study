package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

//015支付
type ZOFPAY struct {
	Notify_url string
	Pay_url    string
	Mer_code   string
	Key        string
	Fxback_url string
}

//015支付
func ZofPay() (int, string, string, string, string, map[string]string) {

	log_path := ""
	api_method := "POST"
	re_status := 100
	re_msg := "请求错误"
	img_url := ""

	api := ZOFPAY{
		Notify_url: "https://www.baidu.com/",
		Pay_url:    "http://www.015pay.com/Pay",
		Mer_code:   "100351",
		Key:        "RaerSVMrHEjQkUelATGZKyiBCSgCSURd",
		Fxback_url: "https://www.baidu.com/",
	}

	p := PayData{
		Amount:       "500.00",
		Order_number: "456287234495823112365",
		Pay_bank:     "ylsm",
		Ip:           "127.0.0.1",
	}

	param_form := map[string]string{
		"fxid":        api.Mer_code,
		"fxdesc":      p.Order_number,
		"fxddh":       p.Order_number,
		"fxnotifyurl": api.Notify_url,
		"fxbackurl":   api.Fxback_url,
		"fxfee":       p.Amount,
		"fxpay":       p.Pay_bank,
		"fxip":        p.Ip,
		"fxuserid":    api.Mer_code,
	}

	//拼接
	result_url := MapCreatLink(param_form, "fxid,fxdesc,fxfee,fxnotifyurl", "", 2)
	result_url += fmt.Sprintf("%s", api.Key)
	sign := common.HexMd5(result_url)
	param_form["fxsign"] = sign

	//请求三方接口
	param := MapCreatLinkSort(param_form, "&", true, false)

	//把post form 表单提交 发送给目标服务器
	resp, err := http.PostForm(api.Pay_url, url.Values{
		"fxid":        {param_form["fxid"]},
		"fxdesc":      {param_form["fxdesc"]},
		"fxddh":       {param_form["fxddh"]},
		"fxnotifyurl": {param_form["fxnotifyurl"]},
		"fxbackurl":   {param_form["fxbackurl"]},
		"fxfee":       {param_form["fxfee"]},
		"fxpay":       {param_form["fxpay"]},
		"fxip":        {param_form["fxip"]},
		"fxuserid":    {param_form["fxuserid"]},
		"fxsign":      {param_form["fxsign"]},
	})
	if err != nil {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	common.LogsWithFileName(log_path, "zofpay_create_", "param->"+param+"\nmsg->"+string(body))
	if err != nil {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	var json_res map[string]interface{}
	err = json.Unmarshal(body, &json_res)
	if err != nil {
		re_msg = "json错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	if fmt.Sprintf("%v", json_res["status"]) != "1" {
		re_msg = fmt.Sprintf("%v", json_res["error"])
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	img_url = fmt.Sprintf("%v", json_res["payurl"])

	if img_url == "" {
		re_msg = "接口错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	re_status = 200
	re_msg = "success"
	return re_status, re_msg, api_method, img_url, img_url, param_form

}

//根据想要的进行拼接
// mid 拼接符号
// dm 为要拼接的内容用，隔开
// url 0为url模式 ; 1 拼接：参数名值mid参数名值; 2 值拼接
func MapCreatLink(m map[string]string, dm string, mid string, url int) string {
	var result = ""
	Extract := strings.Split(dm, ",")
	for _, value := range Extract {
		if value != "" && m[value] != "" {
			switch url {
			case 0:
				result += value + "=" + fmt.Sprintf("%s", m[value]) + mid
			case 1:
				result += value + fmt.Sprintf("%s", m[value]) + mid
			case 2:
				result += fmt.Sprintf("%s", m[value]) + mid
			}
		}
	}

	if mid != "" {
		result = strings.TrimRight(result, mid)
	}

	return result
}
