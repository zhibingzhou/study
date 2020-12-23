package common

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v7"
)

type RedisConf struct {
	Host   string
	Port   string
	Pwd    string
	DBName int
}

var pool *redis.Client

// redis初始化
func InitRedis(redisMsg RedisConf) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     redisMsg.Host + ":" + redisMsg.Port,
		Password: redisMsg.Pwd,
		DB:       redisMsg.DBName,
	})
	err := client.Ping().Err()
	if err != nil {
		log.Fatalln(err)
	}

	return client
}

func example() {
	key := ""
	value := ""
	//查询是否有值
	if !pool.SIsMember(key, value).Val() {

	}

	err := pool.Set("waterMark", "value", -1).Err()
	if err != nil {

	}

	pic := pool.Get("waterMark").Val()
	if pic != "" {

	}

	//删除
	pool.Del(key)
	//删除集合中的一个元素
	pool.SRem(key, value)
	//添加
	pool.SAdd(key, value)

	//也可以获取一个key
	teamA := pool.HGet(key, value).Val()
	if teamA != "" {

	}

	//可以获取集合中某个元素的值
	pool.HGet(key, value).Val()

	//自增
	err = pool.Incr(key).Err()
	if err != nil {
	}

	//redis事务
	pipe := pool.TxPipeline()
	defer pipe.Close()
	pipe.SAdd("key", "value")

	pipe.HSet("key", "name", "value")
    pool.HGet("key", "name").Val()

}

//存多个值
func MatchByNickName(nickName string) map[string]string {

	redisKey := "matchs:nick_name:" + nickName
	//var match Matches
	res := map[string]string{}

	//优先查询redis //存map
	dMap, err := pool.HMGet(redisKey, "id", "nick_name").Result()
	if err != nil || len(dMap) < 0 {

		// 查询数据库

		// _, err := b.Table(db, "matchs").Select(&match, b.Where(b.Eq("nick_name", nickName)), b.Limit(1))
		// if err != nil || match.ID < 1 {
		// 	return res
		// }

		val := map[string]interface{}{}
		val["id"] = ""
		val["nick_name"] = ""
		err = pool.HMSet(redisKey, val).Err()
		if err != nil {
			return res
		}
		
		//pool.HGet(redisKey, id).Val()
		dMap, err = pool.HMGet(redisKey, "id", "nick_name").Result()
		if err != nil {
			return res
		}
	}

	res["id"] = fmt.Sprintf("%v", dMap[0])
	res["nick_name"] = fmt.Sprintf("%v", dMap[1])

	return res

}
