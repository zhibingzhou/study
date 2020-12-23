package common

import (
	"fmt"

	"github.com/axgle/mahonia"
	"github.com/crawlab-team/crawlab-go-sdk/entity"
	"github.com/gocolly/colly/v2"
)

func city() []map[string]interface{} {

	var data []map[string]interface{}
	// 生成 colly 采集器
	c := colly.NewCollector(
		colly.AllowedDomains("www.xiangqinwang.cn"),
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36"),
	)

	// 抓取结果数据钩子函数
	c.OnHTML("dl", func(e *colly.HTMLElement) {

		if e.ChildAttr("dl >dt >a", "href") != "" {

			// 抓取结果实例
			item := entity.Item{
				"title": e.ChildText("dl >dt >a"),
				"url":   "https://www.xiangqinwang.cn" + e.ChildAttr("dl >dt >a", "href"),
			}
			data = append(data, item)

			// 打印抓取结果
			fmt.Println(item)

		}
		// 取消注释调用 Crawlab Go SDK 存入数据库
		//_ = crawlab.SaveItem(item)
	})

	// 访问初始 URL
	startUrl := "https://www.xiangqinwang.cn/map.asp"
	_ = c.Visit(startUrl)

	// 等待爬虫结束
	c.Wait()

	return data
}

func page(Gourl string) []map[string]interface{} {

	var data []map[string]interface{}

	// 生成 colly 采集器
	c := colly.NewCollector(
		colly.AllowedDomains("www.xiangqinwang.cn"),
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36"),
	)

	// 分页钩子函数
	c.OnHTML(".dh > option", func(e *colly.HTMLElement) {

		paged := entity.Item{
			"title": mahonia.NewDecoder("gbk").ConvertString(e.ChildText("option")),
			"url":   "https://www.xiangqinwang.cn" + mahonia.NewDecoder("gbk").ConvertString(e.Attr("value")),
		}
		data = append(data, paged)
		// 打印抓取结果
		fmt.Println("页数", paged["title"], paged["url"])
		//_ = c.Visit("https://www.xiangqinwang.cn")

	})

	// 访问初始 URL
	startUrl := Gourl
	_ = c.Visit(startUrl)

	// 等待爬虫结束
	c.Wait()

	return data
}

func userpage(Gourl string) []map[string]interface{} {

	var data []map[string]interface{}

	// 生成 colly 采集器
	c := colly.NewCollector(
		colly.AllowedDomains("www.xiangqinwang.cn"),
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36"),
	)

	// // 抓取结果数据钩子函数
	c.OnHTML(".oe_userlist_layout > dl", func(e *colly.HTMLElement) {
		// 抓取结果实例
		item := entity.Item{
			"title": e.ChildText("dt > a"),
			"url":   "https://www.xiangqinwang.cn" + e.ChildAttr("dt > a", "href"),
		}
		data = append(data, item)
		// 打印抓取结果
		fmt.Println(item)

		// 取消注释调用 Crawlab Go SDK 存入数据库
		//_ = crawlab.SaveItem(item)
	})

	// 访问初始 URL
	startUrl := Gourl
	_ = c.Visit(startUrl)

	// 等待爬虫结束
	c.Wait()

	return data
}

func merinfo(Gourl string) []map[string]interface{} {

	var data []map[string]interface{}

	// 生成 colly 采集器
	c := colly.NewCollector(
		colly.AllowedDomains("www.xiangqinwang.cn"),
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36"),
	)

	// // 抓取结果数据钩子函数
	c.OnHTML(".new_pd_about", func(e *colly.HTMLElement) {

		// 抓取结果实例
		item := entity.Item{
			"title": mahonia.NewDecoder("gbk").ConvertString(e.ChildText("div.new_pd_about_label > ul")),
		}
		fmt.Println(item)

		// 打印抓取结果
		item = entity.Item{
			"title": mahonia.NewDecoder("gbk").ConvertString(e.ChildText("div.new_pd_about_label2 > ul")),
		}
		fmt.Println(item)

		// 取消注释调用 Crawlab Go SDK 存入数据库
		//_ = crawlab.SaveItem(item)
	})

	// 访问初始 URL
	startUrl := Gourl
	_ = c.Visit(startUrl)

	// 等待爬虫结束
	c.Wait()

	return data
}
