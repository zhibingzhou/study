package model

import (
	b "github.com/orca-zhang/borm"
)

func ListTotalbyBorm(table_name string, p_where []interface{}) (int, error) {
	total := 0
	t := b.Table(bdb.DB, table_name).Debug()
	_, err := t.Select(&total, b.Fields("count(*)"), b.Where(p_where...))
	return total, err
}

func GetInformationborm(table_name, order_by string, page_size, offset int, p_where []interface{}) ([]Information, error) {

	records := []Information{}
	if order_by == "" {
		order_by = "id" + " desc"
	}
	t := b.Table(bdb.DB, table_name).Debug()
	_, err := t.Select(&records, b.Fields("id", "height", "note", "job", "weight", "address", "age", "name", "body", "moneymax", "moneymin", "gourl", "marry", "school", "imgurl", "gender"), b.Where(p_where...), b.OrderBy(order_by), b.Limit(offset, page_size))

	return records, err
}

func GetAllSchoolborm(table_name string, page_size, offset int, p_where []interface{}) ([]School_list, error) {

	order_by := "id" + " desc"

	c_list := []School_list{}
	t := b.Table(bdb.DB, table_name).Debug()

	_, err := t.Select(&c_list, b.Fields("id", "school_title"), b.Where(p_where...), b.OrderBy(order_by), b.Limit(offset, page_size))

	return c_list, err
}

func GetAllBodyborm(table_name string, page_size, offset int, p_where []interface{}) ([]Body_list, error) {

	order_by := "id" + " desc"

	c_list := []Body_list{}
	t := b.Table(bdb.DB, table_name).Debug()

	_, err := t.Select(&c_list, b.Fields("id", "body_title"), b.Where(p_where...), b.OrderBy(order_by), b.Limit(offset, page_size))

	return c_list, err

}

func GetAllMarryborm(table_name string, page_size, offset int, p_where []interface{}) ([]Marry_list, error) {

	order_by := "id" + " desc"

	c_list := []Marry_list{}
	t := b.Table(bdb.DB, table_name).Debug()

	_, err := t.Select(&c_list, b.Fields("id", "marry_title"), b.Where(p_where...), b.OrderBy(order_by), b.Limit(offset, page_size))

	return c_list, err
}

func GetAllJobborm(table_name string, page_size, offset int, p_where []interface{}) ([]Job_list, error) {

	order_by := "id" + " desc"

	c_list := []Job_list{}

	t := b.Table(bdb.DB, table_name).Debug()

	_, err := t.Select(&c_list, b.Fields("id", "job_title"), b.Where(p_where...), b.OrderBy(order_by), b.Limit(offset, page_size))

	return c_list, err
}
