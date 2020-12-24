package common

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisConf struct {
	Host   string
	Port   string
	Pwd    string
	DBName int
}

var pool *redis.Client
var ctx = context.Background()

// redis初始化
func InitRedis(redisMsg RedisConf) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     redisMsg.Host + ":" + redisMsg.Port,
		Password: redisMsg.Pwd,
		DB:       redisMsg.DBName,
	})
	err := client.Ping(ctx).Err()
	if err != nil {
		log.Fatalln(err)
	}

	return client
}

func example() {
	key := ""
	value := ""
	//查询是否有值
	if !pool.SIsMember(ctx, key, value).Val() {

	}

	err := pool.Set(ctx, "waterMark", "value", -1).Err()
	if err != nil {

	}

	//覆盖值 ，如果不存在为设置的值
	// SET key value EX 10 NX
	set, err := pool.SetNX(ctx, "key", "value", 10*time.Second).Result()

	// SET key value keepttl NX
	set, err = pool.SetNX(ctx, "key", "value", redis.KeepTTL).Result()

	fmt.Println(set, err)

	pic := pool.Get(ctx, "waterMark").Val()
	if pic != "" {

	}

	//删除
	pool.Del(ctx, key)
	//删除集合中的一个元素
	pool.SRem(ctx, key, value)
	//添加
	pool.SAdd(ctx, key, value)

	//删除集合中的一个元素
	pool.ZRem(ctx, key)

	//也可以获取一个key
	teamA := pool.HGet(ctx, key, value).Val()
	if teamA != "" {

	}

	//可以获取集合中某个元素的值
	pool.HGet(ctx, key, value).Val()

	//自增
	err = pool.Incr(ctx, key).Err()
	if err != nil {
	}

	//redis事务
	pipe := pool.TxPipeline()
	defer pipe.Close()
	pipe.SAdd(ctx, "key", "value")

	pipe.HSet(ctx, "key", "name", "value")
	pool.HGet(ctx, "key", "name").Val()

}

//存多个值
func MatchByNickName(nickName string) map[string]string {

	redisKey := "matchs:nick_name:" + nickName
	//var match Matches
	res := map[string]string{}

	//优先查询redis //存map
	dMap, err := pool.HGetAll(ctx, redisKey).Result()
	if err != nil || len(dMap) < 0 {

		// 查询数据库

		// _, err := b.Table(db, "matchs").Select(&match, b.Where(b.Eq("nick_name", nickName)), b.Limit(1))
		// if err != nil || match.ID < 1 {
		// 	return res
		// }

		// val := map[string]interface{}{}
		// val["id"] = ""
		// val["nick_name"] = ""
		// err = pool.HMSet(ctx, redisKey, val).Err()
		// if err != nil {
		// 	return res
		// }

		// //pool.HGet(redisKey, id).Val()
		// dMap, err = pool.HMGet(ctx, redisKey, "id", "nick_name").Result()
		// if err != nil {
		// 	return res
		// }
	}

	// res["id"] = fmt.Sprintf("%v", dMap[0])
	// res["nick_name"] = fmt.Sprintf("%v", dMap[1])

	return res

}
