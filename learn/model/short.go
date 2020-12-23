package model

import (
	"time"

	"github.com/zhibingzhou/go_public/common"
	"github.com/zhibingzhou/go_public/redis"
)

type Redirect struct {
	Code      string    `json:"code"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

func Write_url_redis(mer_Info Redirect) int {
	redis_key := "Redirect:" + mer_Info.Code
	//优先查询redis
	a_map := redis.RediGo.HgetAll(redis_key)
	if len(a_map["id"]) < 1 {
		a_map = common.StructToMapSlow(mer_Info)
		redis.RediGo.Hmset(redis_key, a_map, redis_data_time)
		redis.RediGo.Sadd(Data_Redis_Key, redis_key, redis_data_time)
		return 200
	}
	return 100
}

func Read_url_redis(code string) (int, string) {
	status := 100
	redis_key := "Redirect:" + code
	//优先查询redis
	a_map := redis.RediGo.HgetAll(redis_key)
	if len(a_map["id"]) < 1 {
		return status, ""
	}
	return 200, a_map["url"]
}
