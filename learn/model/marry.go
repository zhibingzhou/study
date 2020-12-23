package model

import (

	"github.com/zhibingzhou/go_public/common"
	"github.com/zhibingzhou/go_public/redis"
)


/**
*  根据名称查询
 */
 func MerInfoByMarry(title string) Marry_list {
	var p_list Marry_list
	gdb.DB.Model(&Marry_list{}).Where("marry_title = ?", title).First(&p_list)
	return p_list
}


//去重
func Marry_listRedis(title string, mer_Info Marry_list) map[string]string {
	redis_key := "Marry_list:" + title
	//优先查询redis
	a_map := redis.RediGo.HgetAll(redis_key)
	if len(a_map["id"]) < 1 {
		a_info := MerInfoByMarry(title)
		if a_info.Id < 1 {

			a_map = common.StructToMapSlow(mer_Info)
			mer_sql := InsertSql("marry_list", a_map)
			err := Query(mer_sql)
			if err != nil {
				return a_map
			}

			redis.RediGo.Hmset(redis_key, a_map, redis_data_time)
			redis.RediGo.Sadd(Data_Redis_Key, redis_key, 0)
		}
	}
	return a_map
}