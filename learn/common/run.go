package common

import (
	"fmt"
)

func Run(url string) {

	work := NewChannelPool(10, true)

	work.Run()

	work.WorkGO <- Request{Url: GoUrl + url, ParserFunc: ParseCityList}

	for {
		result := <-work.Result //卡住等待 ParseResult 5.

		for _, item := range result.Items {
			fmt.Println(item)
		}

		for _, request := range result.Requests {
			work.WorkGO <- request
		}
	}

}

func KJ() {

	data := city()
	allusr := []map[string]interface{}{}
	for _, value := range data {
		pagedata := page(fmt.Sprintf("%s", value["url"]))
		for _, valuec := range pagedata {
			allusr = append(allusr, userpage(fmt.Sprintf("%s", valuec["url"]))...)
		}
	}

	for _, value := range allusr {
		merinfo(fmt.Sprintf("%s", value["url"]))
	}

}
