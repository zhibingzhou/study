package test

import (
	"learn/common"
	"testing"

	"github.com/axgle/mahonia"
	//"github.com/axgle/mahonia"
)

func TestGet(t *testing.T) {

	//过程
	body, _ := common.Fetch(common.GoUrl + "/tianjin_hexi/")
	bodystr := mahonia.NewDecoder("gbk").ConvertString(string(body))
	common.WriteToFile(bodystr)
	//结果
	common.PageUserList([]byte(bodystr))

}
