package model
import (

	"github.com/zhibingzhou/go_public/common"
	"github.com/zhibingzhou/go_public/redis"
)

/**
*  根据id查询
 */
 func MerInfoById(id string) Information {
	var p_list Information
	gdb.DB.Model(&Information{}).Where("id = ?", id).First(&p_list)
	return p_list
}




//去重
func MerInfoRedis(mer_id string, mer_Info Information) map[string]string {
	redis_key := "mer_Info:" + mer_id
	//优先查询redis
	a_map := redis.RediGo.HgetAll(redis_key)
	if len(a_map["id"]) < 1 {
		a_info := MerInfoById(mer_id)
		if a_info.Id < 1 {

			a_map = common.StructToMapSlow(mer_Info)
			mer_sql := InsertSql("information", a_map)
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