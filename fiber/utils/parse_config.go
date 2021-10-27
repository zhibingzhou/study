package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// 参数处理
func dbConfig() *DataBase {
	//数据库的配置
	dbFile := "./conf/database.json"
	confByte, err := ioutil.ReadFile(dbFile)
	if err != nil {
		panic(err)
	}
	var dbConf *DataBase
	//解析json格式
	err = json.Unmarshal(confByte, &dbConf)
	if err != nil {
		panic(err)
	}

	return dbConf
}

// 参数处理
func appConfig() *AppConfig {
	//网站的配置
	appFile := "./conf/conf.json"
	confByte, err := ioutil.ReadFile(appFile)
	if err != nil {
		panic(err)
	}
	var appConf *AppConfig
	//解析json格式
	err = json.Unmarshal(confByte, &appConf)
	if err != nil {
		panic(err)
	}

	return appConf
}

/*
* 排序拼接
* mid 拼接符号
* url 是否是url模式
* conectnil 空值是否参与拼接 为true 参与拼接
 */
func MapCreatLinkSort(m map[string]string, mid string, url bool, conectnil bool) string {
	var result = ""

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, k := range keys {
		if conectnil {
			if url {
				result += k + "=" + fmt.Sprintf("%s", m[k]) + mid
			} else {
				result += k + fmt.Sprintf("%s", m[k]) + mid
			}
		} else if m[k] != "" {
			if url {
				result += k + "=" + fmt.Sprintf("%s", m[k]) + mid
			} else {
				result += k + fmt.Sprintf("%s", m[k]) + mid
			}
		}

	}

	if mid != "" {
		result = strings.TrimRight(result, mid)
	}

	return result
}
