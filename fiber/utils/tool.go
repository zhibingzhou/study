package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

/** generate md5 checksum of URL in hex format **/
func HexMd5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	c := m.Sum(nil)
	return hex.EncodeToString(c)
}

func Random(param string, length int) string {
	str := ""
	if length < 1 {
		return str
	}
	tmp := "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	switch param {
	case "number":
		tmp = "1234567890"
	case "small":
		tmp = "abcdefghijklmnopqrstuvwxyz"
	case "big":
		tmp = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case "smallnumber":
		tmp = "1234567890abcdefghijklmnopqrstuvwxyz"
	case "bignumber":
		tmp = "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case "bigsmall":
		tmp = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	leng := len(tmp)
	ran := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		s_ind := ran.Intn(leng)
		str = str + Substr(tmp, s_ind, 1)
	}

	return str
}

/**
*  start：正数 - 在字符串的指定位置开始,超出字符串长度强制把start变为字符串长度
*  负数 - 在从字符串结尾的指定位置开始
*  0 - 在字符串中的第一个字符处开始
*  length:正数 - 从 start 参数所在的位置返回
*  负数 - 从字符串末端返回
 */
func Substr(str string, start, length int) string {
	if length == 0 {
		return ""
	}
	rune_str := []rune(str)
	len_str := len(rune_str)

	if start < 0 {
		start = len_str + start
	}
	if start > len_str {
		start = len_str
	}
	end := start + length
	if end > len_str {
		end = len_str
	}
	if length < 0 {
		end = len_str + length
	}
	if start > end {
		start, end = end, start
	}
	return string(rune_str[start:end])
}

func HttpBody(url, method, param string, header map[string]string) (int, []byte) {
	code := 100
	par := strings.NewReader(param)
	req, err := http.NewRequest(method, url, par)

	if err != nil {
		panic(err)
		return code, nil
	}

	if req.Body != nil {
		defer req.Body.Close()
	}

	if len(header) < 1 {
		header["Content-Type"] = "text/plain; charset=UTF8"
		header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.80 Safari/537.36"
		header["Accept"] = "application/json, text/plain, */*"
	}
	for h_k, h_v := range header {
		req.Header.Set(h_k, h_v)
	}

	tr := &http.Transport{DisableKeepAlives: true,
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*30) //设置建立连接超时
			if err != nil {
				return nil, err
			}
			c.SetDeadline(time.Now().Add(5 * time.Second)) //设置发送接收数据超时
			return c, nil
		}}
	client := &http.Client{Transport: tr}

	//处理返回结果
	resp, err := client.Do(req)
	if err != nil {
		return code, nil
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return code, nil
	}

	code = resp.StatusCode

	return code, msg
}
