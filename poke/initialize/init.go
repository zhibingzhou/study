package initialize

import "poke/global"

func init(){
	Viper()
	global.GVA_DB = GormMysql()
}