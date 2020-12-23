package abb

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"testing"

	"gitlab.stagingvip.net/publicGroup/public/common"
)

func TestGurl(t *testing.T) {

}

func TestAbs(t *testing.T) {
	var a, expect float64 = -10, 10

	actual := math.Abs(a)
	if actual != expect {
		t.Fatalf("a = %f, actual = %f, expected = %f", a, actual, expect)
	}
}

func TestCode(t *testing.T) {
	//content :=  `【柳州银行】您尾号7776账户于12月01日02:04转入500.00元，对方户名：支付宝（中国）网络技术有限公司（摘要：李铁升支付宝转账）活期余额为9992.00元。客服热线0772-9628`
	content := `【招商银行】您账户2878于12月19日他行实时转入人民币10000.00，付方王成。戴森 cmbt.cn/sddj 。  `
	a, b, c := ttt(content)
	fmt.Println(a, b, c)
}

func ttt(content string) (string, string, bool) {
	card_num := ""
	amount := ""
	result := false
	//通过切割，获取尾号信息
	c_arr := strings.Split(content, "账户")
	if len(c_arr) > 1 {
		card_num = common.Substr(c_arr[1], 0, 4)
		result = true
	}
	c_arr1 := strings.Split(content, "转入")
	if len(c_arr1) < 2 {
		c_arr1 = strings.Split(content, "存入")
		if len(c_arr1) < 2 {
			c_arr1 = strings.Split(content, "收款")
			if len(c_arr1) < 2 {
				return amount, card_num, result
			}
			//通过元去切割
			c_arr2 := strings.Split(c_arr1[1], "人民币")
			reg := regexp.MustCompile(`[\d]|[.]`) // 查找连续的汉字
			if len(c_arr2) < 2 {
				return amount, card_num, result
			}
			amount_arr := reg.FindAllString(c_arr2[1], -1)
			if len(amount_arr) < 1 {
				return amount, card_num, result
			}
			for _, a_val := range amount_arr {
				amount = amount + a_val
			}
			return amount, card_num, result
		}
	}
	//通过元去切割
	c_arr2 := strings.Split(c_arr1[1], "，")
	reg := regexp.MustCompile(`[\d]|[.]`) // 查找连续的汉字
	amount_arr := reg.FindAllString(c_arr2[0], -1)
	if len(amount_arr) < 1 {
		return amount, card_num, result
	}
	for _, a_val := range amount_arr {
		amount = amount + a_val
	}
	return amount, card_num, result
}

func DeleteExtraSpace(s string) string {
	//删除字符串中的多余空格，有多个空格时，仅保留一个空格
	s1 := strings.Replace(s, "  ", " ", -1)      //替换tab为空格
	regstr := "\\s{2,}"                          //两个及两个以上空格的正则表达式
	reg, _ := regexp.Compile(regstr)             //编译正则表达式
	s2 := make([]byte, len(s1))                  //定义字符数组切片
	copy(s2, s1)                                 //将字符串复制到切片
	spc_index := reg.FindStringIndex(string(s2)) //在字符串中搜索
	for len(spc_index) > 0 {                     //找到适配项
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) //删除多余空格
		spc_index = reg.FindStringIndex(string(s2))            //继续在字符串中搜索
	}
	return string(s2)
}
