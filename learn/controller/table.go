package controller

import (
	"fmt"
	thread "learn/Thread"
	"learn/model"

	"github.com/gin-gonic/gin"
)

func GetAllData(c *gin.Context) {

	d := map[string]interface{}{}
	m := []model.Information{}
	gender := c.Query("gender")
	body := c.Query("body")
	job := c.Query("job")
	marry := c.Query("marry")
	school := c.Query("school")
	maxmoney := c.Query("maxmoney")
	age := c.Query("age")
	page := c.Query("page")
	page_size := c.Query("page_size")
	c_status := 100
	c_msg := "请求失败"

	c_status, d["total"], c_msg, m = thread.BGetAllData(gender, body, job, marry, school, maxmoney, age, page, page_size)

	total := d["total"].(int)
	pages := total / 20

	_, _, _, schools := thread.BGetAllSchoolborm()

	_, _, _, bodys := thread.BGetAllBodyborm()

	_, _, _, jobs := thread.BGetAllJobborm()

	_, _, _, marrys := thread.BGetAllMarryborm()

	c.HTML(200, "show.tpl", gin.H{
		"c_msg":    c_msg,
		"c_status": c_status,
		"total":    d["total"],
		"school":   schools,
		"body":     bodys,
		"marry":    marrys,
		"job":      jobs,
		"page":     pages,
		"list":     m,
		"title":    "爬虫数据展示",
	})

}

func GoLongUrl(c *gin.Context) {

	code := c.Request.URL

	fmt.Println(code)

	_, pay_url := thread.GOLongUrl("")

	tpl_name := "jump.tpl"

	c.HTML(200, tpl_name, gin.H{
		"pay_url": pay_url,
	})

}

func KeepUrl(c *gin.Context) {

	d := map[string]interface{}{}
	c_msg := "请求失败"
	url := c.PostForm("url")

	c_status, code := thread.Savetoredis(url)

	if c_status == 200 {
		c_msg = SystemUrl + "/" + code
	}

	//将数据装载到json返回值
	c.JSON(200, &JsonOut{Status: c_status, Msg: c_msg, Data: d})
}
