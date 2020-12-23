package main

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/gin-gonic/gin"
// 	"gitlab.stagingvip.net/publicGroup/public/common"
// )

// /**
// * tdpay的回调接口
// * 隐式回调接口，就是服务器对服务器端的推送
//  */
// func Calljpay(c *gin.Context) {

// 	res := "fail"
// 	sign_str := ""
// 	param := map[string]string{
// 		"fxid":          c.PostForm("fxid"),
// 		"fxorderid":     c.PostForm("fxorderid"),
// 		"fxtranid":      c.PostForm("fxtranid"),
// 		"fxamount":      c.PostForm("fxamount"),
// 		"fxamount_succ": c.PostForm("fxamount_succ"),
// 		"fxstatus":      c.PostForm("fxstatus"),
// 		"fxremark":      c.PostForm("fxremark"),
// 		"fxsign":        c.PostForm("fxsign"),
// 	}
// 	sign := param["fxsign"]
// 	if param["status"] == "1" {
// 		sign_str := fmt.Sprintf("%s%s%s%s%s", param["fxstatus"], param["fxid"], param["fxamount_succ"], param["fxorderid"], param["fxamount"])
// 		is_cent := 0
// 		call_status, _ := thread.ThreadUpdatePay(param["fxorderid"], param["fxtranid"], param["fxamount_succ"], sign, sign_str, is_cent)
// 		if call_status == 200 {
// 			res = "success"
// 		}

// 	}

// 	common.LogsWithFileName(log_path, "tdpay_call_", "sign->"+sign+"\nsign_str->"+sign_str+"\nres->"+res)

// 	c.Writer.WriteString(res)
// }

// /**
// * tdpay的回调接口
// * 隐式回调接口，就是服务器对服务器端的推送
//  */
// func Callxpay(c *gin.Context) {

// 	res := "fail"
// 	sign_str := ""
// 	param := map[string]string{
// 		"fxid":          c.PostForm("fxid"),
// 		"fxorderid":     c.PostForm("fxorderid"),
// 		"fxtranid":      c.PostForm("fxtranid"),
// 		"fxamount":      c.PostForm("fxamount"),
// 		"fxamount_succ": c.PostForm("fxamount_succ"),
// 		"fxstatus":      c.PostForm("fxstatus"),
// 		"fxremark":      c.PostForm("fxremark"),
// 		"fxsign":        c.PostForm("fxsign"),
// 	}
// 	sign := param["fxsign"]
// 	if param["status"] == "1" {
// 		sign_str := MapCreatLinkSort(param, "&", true, false)
// 		is_cent := 0
// 		call_status, _ := thread.ThreadUpdatePay(param["fxorderid"], param["fxtranid"], param["fxamount_succ"], sign, sign_str, is_cent)
// 		if call_status == 200 {
// 			res = "success"
// 		}

// 	}

// 	common.LogsWithFileName(log_path, "tdpay_call_", "sign->"+sign+"\nsign_str->"+sign_str+"\nres->"+res)

// 	c.Writer.WriteString(res)
// }

// /**
// * xxfpay的回调接口
// * 隐式回调接口，就是服务器对服务器端的推送
//  */
// func Callxxfpay(c *gin.Context) {

// 	res := "fail"
// 	sign_str := ""
// 	param := map[string]string{
// 		"channel":    c.PostForm("channel"),
// 		"tradeNo":    c.PostForm("tradeNo"),
// 		"outTradeNo": c.PostForm("outTradeNo"),
// 		"money":      c.PostForm("money"),
// 		"realMoney":  c.PostForm("realMoney"),
// 		"uid":        c.PostForm("uid"),
// 		"outUserId":  c.PostForm("outUserId"),
// 		"outBody":    c.PostForm("outBody"),
// 	}
// 	sign := param["sign"]

// 	if param["status"] == "1" {
// 		sign_str, _ := json.Marshal(param)
// 		is_cent := 0
// 		call_status, _ := thread.ThreadUpdatePay(param["outTradeNo"], param["outTradeNo"], param["realMoney"], sign, sign_str, is_cent)
// 		if call_status == 200 {
// 			res = "SUCCESS"
// 		}

// 	}

// 	common.LogsWithFileName(log_path, "xxfpay_call_", "sign->"+sign+"\nsign_str->"+sign_str+"\nres->"+res)

// 	c.Writer.WriteString(res)
// }

// /**
// * zofpay的回调接口
// * 隐式回调接口，就是服务器对服务器端的推送
//  */
// func Callzofpay(c *gin.Context) {

// 	res := "fail"
// 	sign_str := ""
// 	param := map[string]string{
// 		"fxid":     c.PostForm("fxid"),
// 		"fxddh":    c.PostForm("fxddh"),
// 		"fxorder":  c.PostForm("fxorder"),
// 		"fxdesc":   c.PostForm("fxdesc"),
// 		"fxfee":    c.PostForm("fxfee"),
// 		"fxattch":  c.PostForm("fxattch"),
// 		"fxstatus": c.PostForm("fxstatus"),
// 		"fxtime":   c.PostForm("fxtime"),
// 	}
// 	sign := c.PostForm("fxsign")
// 	if param["status"] == "1" {
// 		sign_str := MapCreatLink(param, "fxstatus,fxid,fxddh,fxfee", "", 2)
// 		is_cent := 0
// 		call_status, _ := thread.ThreadUpdatePay(param["fxddh"], param["fxorder"], param["fxfee"], sign, sign_str, is_cent)
// 		if call_status == 200 {
// 			res = "success"
// 		}

// 	}

// 	common.LogsWithFileName(log_path, "zofpay_call_", "sign->"+sign+"\nsign_str->"+sign_str+"\nres->"+res)

// 	c.Writer.WriteString(res)
// }
