package model

import (
	"fmt"
	"studyfiber/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

type Gorm struct {
	DB *gorm.DB
}

type CountTotal struct {
	Total float64
	Num   int
}

func init() {
	ReloadConf()
}

var Gdb Gorm

func ReloadConf() {
	connStr := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8", utils.DBConf.User, utils.DBConf.Pwd, utils.DBConf.Network, utils.DBConf.Host, utils.DBConf.Port, utils.DBConf.DbName)
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		fmt.Println("connStr->", connStr)
		panic(err)
	}

	lifeTime, _ := time.ParseDuration(utils.DBConf.LifeTime)
	//最大生命周期
	db.DB().SetConnMaxLifetime(lifeTime)
	//连接池的最大打开连接数
	db.DB().SetMaxOpenConns(utils.DBConf.MaxOpen)
	//连接池的最大空闲连接数
	db.DB().SetMaxIdleConns(utils.DBConf.MaxIdle)
	db.SingularTable(true)
	//启用Logger，显示详细日志
	db.LogMode(utils.DBConf.ShowLog)

	// 禁用日志记录器，不显示任何日志
	//db.LogMode(false)
	Gdb = Gorm{DB: db}
}

func Query(sqlStr string) error {
	res := Gdb.DB.Exec(sqlStr)
	return res.Error
}
