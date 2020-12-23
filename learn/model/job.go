package model

import (

	"github.com/zhibingzhou/go_public/common"
	"github.com/zhibingzhou/go_public/redis"
)

/**
*  根据名称查询
 */
 func MerInfoByJob(title string) Job_list {
	var p_list Job_list
	gdb.DB.Model(&Job_list{}).Where("job_title = ?", title).First(&p_list)
	return p_list
}



//去重
func Job_listRedis(title string, mer_Info Job_list) map[string]string {
	redis_key := "Job_list:" + title
	//优先查询redis
	a_map := redis.RediGo.HgetAll(redis_key)
	if len(a_map["id"]) < 1 {
		a_info := MerInfoByJob(title)
		if a_info.Id < 1 {

			a_map = common.StructToMapSlow(mer_Info)
			mer_sql := InsertSql("job_list", a_map)
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