package model

import (
	"github.com/zhibingzhou/go_public/common"
	"github.com/zhibingzhou/go_public/redis"
)

/**
*  根据名称查询
 */
 func MerInfoBybody(title string) Body_list {
	var p_list Body_list
	gdb.DB.Model(&Body_list{}).Where("body_title = ?", title).First(&p_list)
	return p_list
}


//去重
func Body_listRedis(title string, mer_Info Body_list) map[string]string {
	redis_key := "Body_list:" + title
	//优先查询redis
	a_map := redis.RediGo.HgetAll(redis_key)
	if len(a_map["id"]) < 1 {
		a_info := MerInfoBybody(title)
		if a_info.Id < 1 {

			a_map = common.StructToMapSlow(mer_Info)
			mer_sql := InsertSql("body_list", a_map)
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