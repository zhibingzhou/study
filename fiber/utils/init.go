package utils

// mysql属性定义
type DataBase struct {
	Network  string `json:"network"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	LifeTime string `json:"life_time"`
	MaxOpen  int    `json:"max_open"`
	MaxIdle  int    `json:"max_idle"`
	ShowLog  bool   `json:"show_log"`
}

// 系统参数
type AppConfig struct {
	Path       string `json:"path"`
	Tianxin    string `json:"tianxin"`
	PrivateKey string `json:"privateKey"`
	Iv         string `json:"iv"`
	Proid      string `json:"proid"`
	LogPath    string `json:"logpath"`
}

var AppConf *AppConfig
var DBConf *DataBase

func init() {
	DBConf = dbConfig()
	AppConf = appConfig()
}
