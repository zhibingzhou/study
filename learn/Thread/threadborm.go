package thread

import (
	"fmt"
	"learn/model"

	b "github.com/orca-zhang/borm"
)

func BGetAllData(gender, body, job, marry, school, maxmoney, age, page, page_size string) (int, int, string, []model.Information) {
	t_status := 200
	t_msg := "success"
	c_list := []model.Information{}
	table_name := "information"

	condes := []interface{}{}

	if gender != "" {
		condes = append(condes, b.Cond("gender = ?", gender))
	}

	if body != "" {
		condes = append(condes, b.Cond("body = ?", body))
	}

	if job != "" {
		condes = append(condes, b.Cond("job = ?", job))
	}

	if marry != "" {
		condes = append(condes, b.Cond("marry = ?", marry))
	}

	if school != "" {
		condes = append(condes, b.Cond("school = ?", school))
	}

	if maxmoney == "" {
		maxmoney = "0"
	}

	if age == "" {
		age = "200"
	}

	if len(condes) == 0 {
		condes = append(condes, b.Cond("1=1"))
	}

	total, _ := model.ListTotalbyBorm(table_name, condes)

	page_int, size_int := ThreadPage(page, page_size)
	offset := (page_int - 1) * size_int

	c_list, err := model.GetInformationborm(table_name, "", size_int, offset, condes)

	if err != nil {
		fmt.Println(err)
	}

	return t_status, total, t_msg, c_list
}

func BGetAllSchool() (int, int, string, []map[string]string) {

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

func BGetAllBody() (int, int, string, []map[string]string) {

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

func BGetAllMarry() (int, int, string, []map[string]string) {

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

func BGetAllJob() (int, int, string, []map[string]string) {

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

func BGetAllSchoolborm() (int, int, string, []model.School_list) {

	t_status := 200
	t_msg := "success"
	total := 0
	c_list := []model.School_list{}

	table_name := "school_list"

	condes := []interface{}{}
	condes = append(condes, b.Cond("1=1"))
	total, _ = model.ListTotalbyBorm(table_name, condes)

	page_int, size_int := ThreadPage("1", "100")
	offset := (page_int - 1) * size_int

	c_list, _ = model.GetAllSchoolborm(table_name, size_int, offset, condes)

	return t_status, total, t_msg, c_list
}

func BGetAllBodyborm() (int, int, string, []model.Body_list) {

	t_status := 200
	t_msg := "success"
	total := 0
	table_name := "body_list"
	c_list := []model.Body_list{}

	condes := []interface{}{}
	condes = append(condes, b.Cond("1=1"))

	total, _ = model.ListTotalbyBorm(table_name, condes)

	page_int, size_int := ThreadPage("1", "100")
	offset := (page_int - 1) * size_int

	c_list, _ = model.GetAllBodyborm(table_name, size_int, offset, condes)

	return t_status, total, t_msg, c_list
}

func BGetAllMarryborm() (int, int, string, []model.Marry_list) {

	t_status := 200
	t_msg := "success"
	total := 0

	table_name := "marry_list"

	c_list := []model.Marry_list{}

	condes := []interface{}{}
	condes = append(condes, b.Cond("1=1"))
	total, _ = model.ListTotalbyBorm(table_name, condes)

	page_int, size_int := ThreadPage("1", "100")
	offset := (page_int - 1) * size_int

	c_list, _ = model.GetAllMarryborm(table_name, size_int, offset, condes)

	return t_status, total, t_msg, c_list
}

func BGetAllJobborm() (int, int, string, []model.Job_list) {

	t_status := 200
	t_msg := "success"
	total := 0

	c_list := []model.Job_list{}

	table_name := "job_list"

	condes := []interface{}{}
	condes = append(condes, b.Cond("1=1"))
	total, _ = model.ListTotalbyBorm(table_name, condes)

	page_int, size_int := ThreadPage("1", "100")
	offset := (page_int - 1) * size_int

	c_list, _ = model.GetAllJobborm(table_name, size_int, offset, condes)

	return t_status, total, t_msg, c_list
}
