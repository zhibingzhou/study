package common

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func WriteToFile(content string) error {
	f, err3 := os.Create("./output3.txt") //创建文件
	if err3 != nil {
		return err3
	}
	defer f.Close()
	_, err3 = f.WriteString(content) //写入文件(字节数组)
	if err3 != nil {
		return err3
	}
	f.Sync()
	return err3
}

/**
*  文件内容读取,返回二进制数组
 */
func ReadFile(file_pth string) ([]byte, error) {
	f, err := os.Open(file_pth)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
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
