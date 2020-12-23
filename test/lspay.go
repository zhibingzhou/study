package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

//新万顺支付
type LSPAY struct {
	Notify_url   string `json:"notify_url"`
	Call_backurl string `json:"call_backurl"`
	Pay_url      string `json:"pay_url"`
	Mer_code     string `json:"mer_code"`
	Key          string `json:"key"`
}

func lspay() (int, string, string, string, string, map[string]string) {

	var s_format string = "2006-01-02 15:04:05"
	api := LSPAY{
		Notify_url:   "http://api.sfwage.com/call/lspay.do",
		Call_backurl: "http://api.sfwage.com/public/success.do",
		Pay_url:      "http://www.idsee.cn",
		Mer_code:     "86264537",
		Key:          "ZHQhMgitfQWw",
	}

	pay_config, _ := json.Marshal(api)

	fmt.Println(string(pay_config))

	p := PayData{
		Amount:       "300.00",
		Order_number: "4563412aq1a123903c1162365",
		Pay_bank:     "924",
		Ip:           "127.0.0.1",
	}

	log_path := ""
	api_method := "POST"
	re_status := 100
	re_msg := "请求错误"
	img_url := ""

	param_form := map[string]string{
		"pay_amount":      p.Amount,
		"pay_orderid":     p.Order_number,
		"pay_notifyurl":   api.Notify_url,
		"pay_memberid":    api.Mer_code,
		"pay_callbackurl": api.Call_backurl,
		"pay_applydate":   fmt.Sprintf(time.Now().Format(s_format)), //时间,
		"pay_bankcode":    p.Pay_bank,
	}
	img_url = fmt.Sprintf("%s/pay/index", api.Pay_url)
	//拼接

	result_url := common.MapCreatLinkSort(param_form, "&", true, false)
	result_url += fmt.Sprintf("&key=%s", api.Key)
	sign := common.HexMd5(result_url)
	sign = strings.ToUpper(sign)
	param_form["pay_md5sign"] = sign
	param_form["pay_productname"] = "Test"

	param := common.MapCreatLinkSort(param_form, "&", true, false)

	common.LogsWithFileName(log_path, "lspay_create_", "param->"+param+"\nurl->"+img_url)

	sbForm := "<html><head></head><body onload=\"document.forms[0].submit()\"><form name=\"order\" method=\"post\" action=\"PayBuildDomain\">"
	sbForm = strings.Replace(sbForm, "PayBuildDomain", img_url, -1)
	for item := range param_form {
		sbForm += "<input name=\"" + item + "\" type=\"hidden\" value=\"" + param_form[item] + "\" />"
	}
	sbForm += "</form></body></html>"

	fmt.Println(sbForm)

	fmt.Println(img_url)
	re_status = 200
	re_msg = "success"
	return re_status, re_msg, api_method, img_url, img_url, param_form
}

// /**
// * lspay的回调接口
// * 隐式回调接口，就是服务器对服务器端的推送
//  */
// func Calllspay(c *gin.Context) {

// 	res := "fail"
// 	sign_str := ""
// 	param := map[string]string{
// 		"memberid":    c.PostForm("memberid"),
// 		"orderid":        c.PostForm("orderid"),
// 		"amount":     c.PostForm("amount"),
// 		"transaction_id":  c.PostForm("transaction_id"),
// 		"attach":    c.PostForm("attach"),
// 		"datetime": c.PostForm("datetime"),
// 		"returncode":       c.PostForm("returncode"),
// 	}
// 	sign := c.PostForm("sign")
// 	if param["returncode"] == "00" {
// 		sign_str := common.MapCreatLinkSort(param, "&", true, false)
// 		is_cent := 0
// 		call_status, _ := thread.ThreadUpdatePay(param["out_trade_no"], param["out_trade_no"], param["amount_true"], sign, sign_str, is_cent)
// 		if call_status == 200 {
// 			res = "success"
// 		}

// 	}

// 	common.LogsWithFileName(log_path, "cwspay_call_", "sign->"+sign+"\nsign_str->"+sign_str+"\nres->"+res)

// 	c.Writer.WriteString(res)
// }
