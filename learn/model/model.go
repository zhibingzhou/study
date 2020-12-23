package model

import (
	"database/sql"
	"fmt"
	"strings"
)

type Information struct {
	Id       int     `json:"id" borm:"id"`             //ID
	Gender   string  `json:"gender" borm:"gender"`     //男女
	Name     string  `json:"name" borm:"name"`         //网名
	Age      int     `json:"age" borm:"age"`           //年龄
	Job      string  `json:"job" borm:"job"`           //职业
	Height   float64 `json:"height" borm:"height"`     //身高
	Weight   float64 `json:"weight" borm:"weight"`     //体重
	Body     string  `json:"body" borm:"body"`         //身材
	Marry    string  `json:"marry" borm:"marry"`       //婚姻
	School   string  `json:"school" borm:"school"`     //学校
	MoneyMax int     `json:"moneymax" borm:"moneymax"` // 最大收入
	MoneyMin int     `json:"moneymin" borm:"moneymin"` // 最小收入
	Note     string  `json:"note" borm:"note"`         //备注
	GoUrl    string  `json:"gourl" borm:"gourl"`       //链接
	ImgUrl   string  `json:"imgurl" borm:"imgurl"`     //照片
	Address  string  `json:"address" borm:"address"`   //地方
}

type Job_list struct {
	Id        int    `borm:"id"`        //ID
	Job_title string `borm:"job_title"` //
}

type Body_list struct {
	Id         int    `borm:"id"`         //ID
	Body_title string `borm:"body_title"` //
}

type Marry_list struct {
	Id          int    `borm:"id"`          //ID
	Marry_title string `borm:"marry_title"` //
}

type School_list struct {
	Id           int    `borm:"id"`           //ID
	School_title string `borm:"school_title"` //
}

type CountTotal struct {
	Total float64
	Num   int
}

/**
*  生成插入的sql语句
 */
func InsertSql(table_name string, data map[string]string) string {
	sql := ""
	if table_name == "" || len(data) < 1 {
		return sql
	}

	key_str := ","
	val_str := ","
	for d_k, d_v := range data {
		key_str = key_str + ",`" + d_k + "`"
		val_str = val_str + `,"` + d_v + `"`
	}

	key_str = strings.Replace(key_str, ",,", "", 1)
	val_str = strings.Replace(val_str, ",,", "", 1)
	sql = "insert into `" + table_name + "` (" + key_str + ") VALUES (" + val_str + ");"
	return sql
}

func Query(sql_str string) error {
	res := gdb.DB.Exec(sql_str)
	return res.Error
}

func ListTotal(table_name, field string, p_where map[string]interface{}) (int, float64) {
	var c_total CountTotal
	gdb.DB.Table(table_name).Select(field).Where(p_where).Scan(&c_total)

	return c_total.Num, c_total.Total
}

func ListTotalM(table_name, field, money, age string, p_where map[string]interface{}) (int, float64) {
	var c_total CountTotal
	gdb.DB.Table(table_name).Select(field).Where("moneyMax >= ? and age < ? ", money, age).Scan(&c_total)

	return c_total.Num, c_total.Total
}

func PageListM(table_name, order_by string, page_size, offset int, fields []string, money, age string, p_where map[string]interface{}) ([]map[string]string, error) {
	records := []map[string]string{}
	if order_by == "" {
		order_by = fields[0] + " desc"
	}
	u_rows, err := gdb.DB.Table(table_name).Select(fields).Where(p_where).Where("moneyMax >= ? and age < ? ", money, age).Limit(page_size).Order(order_by).Offset(offset).Rows()
	if err != nil {
		return records, err
	}
	records, err = rows2Map(fields, u_rows)
	return records, err
}

func PageList(table_name, order_by string, page_size, offset int, fields []string, p_where map[string]interface{}) ([]map[string]string, error) {
	records := []map[string]string{}
	if order_by == "" {
		order_by = fields[0] + " desc"
	}
	u_rows, err := gdb.DB.Table(table_name).Select(fields).Where(p_where).Limit(page_size).Order(order_by).Offset(offset).Rows()
	if err != nil {
		return records, err
	}
	records, err = rows2Map(fields, u_rows)
	return records, err
}

/*
数据库数据转MAP
*/
func rows2Map(fields []string, u_rows *sql.Rows) ([]map[string]string, error) {

	records := []map[string]string{}
	var err error

	//创建有效切片
	values := make([]interface{}, len(fields))
	//行扫描，必须复制到这样切片的内存地址中去
	scanArgs := make([]interface{}, len(fields))

	for j := range values {
		scanArgs[j] = &values[j]
	}

	for u_rows.Next() {
		err = u_rows.Scan(scanArgs...)

		if err != nil {
			break
		}
		record := map[string]string{}
		for i, col := range values {
			if col == nil {
				record[fields[i]] = ""
				continue
			}
			col_s, ok := col.([]byte)
			if ok {
				record[fields[i]] = string(col_s)
			} else {
				record[fields[i]] = fmt.Sprintf("%v", col)
			}
		}
		records = append(records[0:], record)
	}
	return records, err
}
