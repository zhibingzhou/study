package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

//欧付宝支付
type OFBPAY struct {
	Return_url string
	Pay_url    string
	Source_url string
	Mer_code   string
	Key        string
}

//欧付宝
func oufubaoPay() (int, string, string, string, string, map[string]string) {

	api := OFBPAY{
		Return_url: "https://www.baidu.com/",
		Pay_url:    "http://154.93.139.58:8036/pay/deposit_order.do",
		Mer_code:   "TD0321",
		Key:        "EFREHTYRFDEWSAQ1", //w
		Source_url: "https://www.baidu.com/",
	}

	p := PayData{
		Amount:       "500.00",
		Order_number: "45628749123332342365",
		Pay_bank:     "1", //支付宝扫码
		Ip:           "127.0.0.1",
	}

	log_path := ""
	api_method := "POST"
	re_status := 100
	re_msg := "请求错误"
	img_url := ""

	//请求参数
	param_form := map[string]string{
		"access_code": api.Mer_code,
	}

	//拼接
	param := fmt.Sprintf("AppliedBankId=%s&CustomerName=%s&amount=%s&webOrderNumber=%s&source_url=%s&return_params=%s&return_url=%s",
		p.Pay_bank, api.Mer_code, p.Amount, p.Order_number, api.Source_url, "123", api.Return_url)
	fmt.Println(param)

	param_form["params"], _ = AesEncrypt(param, []byte(api.Key), []byte(api.Key))

	rep := MapCreatLinkSort(param_form, "&", true, false)
	fmt.Println(rep)
	msg_b, h_status := httpPostForm(api.Pay_url, param_form)
	common.LogsWithFileName(log_path, "ofbpay_create_", "param->"+param+"\nmsg->"+string(msg_b))
	if h_status != nil {
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	var json_res map[string]interface{}
	err := json.Unmarshal([]byte(msg_b), &json_res)
	if err != nil {
		re_msg = "json错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	fmt.Println(string(msg_b))

	if fmt.Sprintf("%v", json_res["Status"]) != "200" {
		re_msg = fmt.Sprintf("%v", json_res["Msg"])
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	img_url = fmt.Sprintf("%v", json_res["Data"])

	if img_url == "" {
		re_msg = "接口错误"
		return re_status, re_msg, api_method, img_url, img_url, param_form
	}

	fmt.Println(img_url)

	re_status = 200
	re_msg = "success"
	return re_status, re_msg, api_method, img_url, img_url, param_form
}

//aes cbc 加密
func AesEncrypt(encodeStr string, key []byte, iv []byte) (string, error) {
	encodeBytes := []byte(encodeStr)
	//根据key 生成密文
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encodeBytes = PKCS5Padding(encodeBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func AesDecrypt(decodeStr string, key []byte, iv []byte) ([]byte, error) {
	//先解密base64
	decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(origData, decodeBytes)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func httpPostForm(Url string, strm map[string]string) (string, error) {

	resp, err := http.PostForm(Url,
		url.Values{"access_code": {strm["access_code"]}, "params": {strm["params"]}})
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return "", err
	}
	return fmt.Sprintf(string(body)), nil

}
