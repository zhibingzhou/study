package controller

import (
	"encoding/json"
	"learn/common"
)

/*
* 定义一个json的返回值类型(首字母必须大写)
 */
type JsonOut struct {
	Status int
	Msg    string
	Data   map[string]interface{}
}

var SystemUrl = "127.0.0.1"

func init() {
	conf_byte, err := common.ReadFile("./conf/conf.json")

	if err != nil {
		panic(err)
	}
	var json_conf map[string]string
	//解析json格式
	err = json.Unmarshal(conf_byte, &json_conf)
	if err != nil {
		panic(err)
	}
	SystemUrl = json_conf["url"] + json_conf["port"]
}
