package test

import (
	"context"
	"fmt"
	"learn/common"
	"testing"

	"github.com/axgle/mahonia"
	"github.com/go-redis/redis/v8"
	//"github.com/axgle/mahonia"
)

func TestGet(t *testing.T) {

	//过程
	body, _ := common.Fetch(common.GoUrl + "/tianjin_hexi/")
	bodystr := mahonia.NewDecoder("gbk").ConvertString(string(body))
	common.WriteToFile(bodystr)
	//结果
	common.PageUserList([]byte(bodystr))

}

type People struct {
	name string
	age  int
}

func TestRedis(t *testing.T) {

	pool := common.InitRedis(common.RedisConf{
		Host:   "127.0.0.1",
		Port:   "6379",
		Pwd:    "foobared",
		DBName: 0,
	})
	var ctx = context.Background()

	//map 存储
	val := map[string]interface{}{}
	val["id"] = "abc"
	val["nick_name"] = 123

	flag, err := pool.HMSet(ctx, "im_key", val).Result()
	fmt.Println(flag)

	//拿出map中的数据
	dMap, err := pool.HGetAll(ctx, "im_key2").Result()
	fmt.Println(dMap, err)

	

	//拿出有序集合中的0-50的元素
	valsd, err := pool.ZRange(ctx, "ccc", 0, 50).Result()

	fmt.Println(valsd, err)

	// 添加有序集合 插入成功为1 插入失败为0
	value, err := pool.ZAdd(ctx, "ccc", &redis.Z{Score: 10, Member: "abc"}).Result()
	fmt.Println(value, err)

	//设置最大和最小值  返回有序集合的所有元素和分数
	vals, err := pool.ZRangeByScoreWithScores(ctx, "ccc", &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()

	fmt.Println(vals, err)

	//返回两个集合相同值数量  zadd ccc 88 "nima"  相同的“nima” 的集合数量
	// ZINTERSTORE out 2 zset1 zset2 WEIGHTS 2 3 AGGREGATE SUM
	valss, err := pool.ZInterStore(ctx, "out", &redis.ZStore{
		Keys:    []string{"ccc", "ddd"},
		Weights: []float64{0, 1},
	}).Result()

	fmt.Println(valss, err)

}