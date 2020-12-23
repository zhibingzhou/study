package thread

import (
	"learn/model"
	"strconv"
)

func GetAllData(gender, body, job, marry, school, maxmoney, age, page, page_size string) (int, int, string, []map[string]string) {
	t_status := 200
	t_msg := "success"
	c_list := []map[string]string{}
	count_field := "count(0) as num"
	table_name := "information"
	p_where := map[string]interface{}{}

	if gender != "" {
		p_where["gender"] = gender
	}

	if body != "" {
		p_where["body"] = body
	}

	if job != "" {
		p_where["job"] = job
	}

	if marry != "" {
		p_where["marry"] = marry
	}

	if school != "" {
		p_where["school"] = school
	}

	if maxmoney == "" {
		maxmoney = "0"
	}

	if age == "" {
		age = "200"
	}

	total, _ := model.ListTotalM(table_name, count_field, maxmoney, age, p_where)

	fields := []string{"id", "height", "note", "job", "weight", "address", "age", "name", "body", "moneymax", "moneymin", "gourl", "marry", "school", "imgurl", "gender"}

	page_int, size_int := ThreadPage(page, page_size)
	offset := (page_int - 1) * size_int
	c_list, _ = model.PageListM(table_name, "", size_int, offset, fields, maxmoney, age, p_where)

	return t_status, total, t_msg, c_list
}

func GetAllSchool() (int, int, string, []map[string]string) {

	t_status := 200
	t_msg := "success"
	total := 0
	c_list := []map[string]string{}

	count_field := "count(0) as num"
	table_name := "school_list"
	p_where := map[string]interface{}{}

	total, _ = model.ListTotal(table_name, count_field, p_where)

	fields := []string{"id", "school_title"}

	page_int, size_int := ThreadPage("1", "100")
	offset := (page_int - 1) * size_int
	c_list, _ = model.PageList(table_name, "", size_int, offset, fields, p_where)

	return t_status, total, t_msg, c_list
}

func GetAllBody() (int, int, string, []map[string]string) {

	t_status := 200
	t_msg := "success"
	total := 0
	c_list := []map[string]string{}

	count_field := "count(0) as num"
	table_name := "body_list"
	p_where := map[string]interface{}{}

	total, _ = model.ListTotal(table_name, count_field, p_where)

	fields := []string{"id", "body_title"}

	page_int, size_int := ThreadPage("1", "100")
	offset := (page_int - 1) * size_int
	c_list, _ = model.PageList(table_name, "", size_int, offset, fields, p_where)

	return t_status, total, t_msg, c_list
}

func GetAllMarry() (int, int, string, []map[string]string) {

	t_status := 200
	t_msg := "success"
	total := 0
	c_list := []map[string]string{}

	count_field := "count(0) as num"
	table_name := "marry_list"
	p_where := map[string]interface{}{}

	total, _ = model.ListTotal(table_name, count_field, p_where)

	fields := []string{"id", "marry_title"}

	page_int, size_int := ThreadPage("1", "100")
	offset := (page_int - 1) * size_int
	c_list, _ = model.PageList(table_name, "", size_int, offset, fields, p_where)

	return t_status, total, t_msg, c_list
}

func GetAllJob() (int, int, string, []map[string]string) {

	t_status := 200
	t_msg := "success"
	total := 0
	c_list := []map[string]string{}

	count_field := "count(0) as num"
	table_name := "job_list"
	p_where := map[string]interface{}{}

	total, _ = model.ListTotal(table_name, count_field, p_where)

	fields := []string{"id", "job_title"}

	page_int, size_int := ThreadPage("1", "100")
	offset := (page_int - 1) * size_int
	c_list, _ = model.PageList(table_name, "", size_int, offset, fields, p_where)

	return t_status, total, t_msg, c_list
}

/**
*  处理分页
 */
func ThreadPage(page, page_size string) (int, int) {
	page_int, _ := strconv.Atoi(page)
	if page_int < 1 {
		page_int = 1
	}
	size_int, _ := strconv.Atoi(page_size)
	if size_int < 1 {
		size_int = 20
	} else if size_int > 100 {
		size_int = 100
	}
	return page_int, size_int
}
