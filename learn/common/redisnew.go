package common

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisConf struct {
	Host   string
	Port   string
	Pwd    string
	DBName int
}

var head = "all_key"
var wxhead = "all_key_wx"
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
	pool = client
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
	//删除集合中的一个元素的一个值
	pool.SRem(ctx, key, value)
	//集合中添加
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
	_, err = pipe.Exec(ctx)

	pool.HGet(ctx, "key", "name").Val()

}

//存map案例
func MatchByNickName(nickName string) (map[string]string, error) {

	redisKey := "matchs:nick_name:" + nickName
	//优先查询redis 拿map
	dMap, err := pool.HGetAll(ctx, redisKey).Result()
	if err == nil && len(dMap["id"]) < 1 {
		// 查询数据库 得 map
		val := map[string]interface{}{}
		val["id"] = "1"
		val["nick_name"] = nickName

		err = pool.HMSet(ctx, redisKey, val).Err()
		if err != nil {
			return dMap, err
		}

		//新增无序集合 所有的key头存在无序集合里面
		err = pool.SAdd(ctx, head, redisKey).Err()
		if err != nil {
			return dMap, err
		}

		dMap, err = pool.HGetAll(ctx, redisKey).Result()
		if err != nil {
			return dMap, err
		}
	}

	return dMap, err
}

//清理缓存
func Delcash() error {

	//拿到key头在集合中的数量
	num, err := pool.SCard(ctx, head).Result()
	if err != nil {
		return err
	}
	var i int64
	for i = 0; i < num; i++ {

		//删除一条数据返回被删除的元素，逐个删除，但这个会返回对应元素
		red_key, err := pool.SPop(ctx, head).Result()

		if err != nil {
			return err
		}

		if pool.Del(ctx, red_key).Err() != nil {
			return err
		}

	}

	return err
}

//爬虫去重 ，添加的是有序集合
func PutData(values string) {

	// 添加有序集合 插入成功为1 插入失败为0
	value, err := pool.ZAdd(ctx, "ccc", &redis.Z{Score: 10, Member: values}).Result()
	fmt.Println(value, err)

	if value == 1 { //说明没有这个key
		onlyid := wxhead + GetKey(16)
		//存对应的data到有序集合
		value, err := pool.ZAdd(ctx, "get_data", &redis.Z{Score: 10, Member: onlyid}).Result()
		fmt.Println(value, err)
		ma := map[string]interface{}{}
		ma["key"] = values
		ma["value"] = "123"
		//再存入map参数
		err = pool.HMSet(ctx, onlyid, ma).Err()
		fmt.Println(err)
	}

}

func Getdata() {

	//设置最大和最小值  返回有序集合的所有元素和分数
	vals, err := pool.ZRangeByScoreWithScores(ctx, "get_data", &redis.ZRangeBy{
		Min: "0",
		Max: "50",
	}).Result()

	for _, value := range vals {
		key := value.Member.(string)
		dMap, err := pool.HGetAll(ctx, key).Result()
		fmt.Println(dMap, err)
		pool.Del(ctx, key).Err()
		//删除集合中的一个指定元素
		pool.ZRem(ctx, "get_data", key)
	}

	fmt.Println(vals, err)
}

func GetKey(length int) string {
	sec := strconv.FormatInt(time.Now().Unix(), 10)
	redKey := "model_get_key:" + sec
	randLen := length
	exTime := 1
	preId := ""

	if length > 10 {
		randLen = length - 10
		preId = sec
	}
	randStr := ""
	for i := 0; i < 50; i++ {
		randStr = Random("smallnumber", randLen)
		//新增无序集合 所有的key头存在无序集合里面
		res, err := pool.SAdd(ctx, redKey, randStr, exTime).Result()
		if err == nil && res > 0 {
			break
		}
	}

	keyStr := preId + randStr
	return keyStr
}

func Random(param string, length int) string {
	str := ""
	if length < 1 {
		return str
	}
	tmp := "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	switch param {
	case "number":
		tmp = "1234567890"
	case "small":
		tmp = "abcdefghijklmnopqrstuvwxyz"
	case "big":
		tmp = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case "smallnumber":
		tmp = "1234567890abcdefghijklmnopqrstuvwxyz"
	case "bignumber":
		tmp = "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case "bigsmall":
		tmp = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	leng := len(tmp)
	ran := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		s_ind := ran.Intn(leng)
		str = str + Substr(tmp, s_ind, 1)
	}

	return str
}

/**
*  start：正数 - 在字符串的指定位置开始,超出字符串长度强制把start变为字符串长度
*  负数 - 在从字符串结尾的指定位置开始
*  0 - 在字符串中的第一个字符处开始
*  length:正数 - 从 start 参数所在的位置返回
*  负数 - 从字符串末端返回
 */
func Substr(str string, start, length int) string {
	if length == 0 {
		return ""
	}
	rune_str := []rune(str)
	len_str := len(rune_str)

	if start < 0 {
		start = len_str + start
	}
	if start > len_str {
		start = len_str
	}
	end := start + length
	if end > len_str {
		end = len_str
	}
	if length < 0 {
		end = len_str + length
	}
	if start > end {
		start, end = end, start
	}
	return string(rune_str[start:end])
}
