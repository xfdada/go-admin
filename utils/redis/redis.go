package redis

import (
	"fmt"
	"go-admin/utils/loggers"
	"time"

	"github.com/go-redis/redis"
)

var Redis *redis.Client

func init() {
	NewClient()
}

func NewClient() *redis.Client {
	if Redis != nil {
		return Redis
	}

	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := Redis.Ping().Result()
	if err != nil {
		fmt.Printf("redis connection failed: %v\n", err.Error())
	}

	return Redis
}

// Set 设置键值和存活时间
func Set(key string, value interface{}, time time.Duration) error {
	var err error
	if time > 0 {
		err = Redis.Set(key, value, time*1000*1000*1000).Err()
	} else {
		err = Redis.Set(key, value, 0).Err()
	}
	if err != nil {
		loggers.Logs(fmt.Sprint("RedisSet Error! key:", key, "Details:", err.Error()))
		return err
	}
	return nil
}

// Get 获取值
func Get(key string) (string, error) {
	value, err := Redis.Get(key).Result()
	if err != nil {
		return "", nil
	}
	return value, nil
}

// Exists 查询键是否存在
func Exists(key string) (int64, error) {
	ok, err := Redis.Exists(key).Result()
	return ok, err
}

func GetResult(key string) (interface{}, error) {
	v, err := Redis.Do("GET", key).Result()
	if err == redis.Nil {
		return v, nil
	}
	return v, err
}

func GetInt(key string) (int, error) {
	v, err := Redis.Do("GET", key).Int()
	if err == redis.Nil {
		return 0, nil
	}
	return v, err
}

func GetInt64(key string) (int64, error) {
	v, err := Redis.Do("GET", key).Int64()
	if err == redis.Nil {
		return 0, nil
	}
	return v, err
}

func RedisGetUint64(key string) (uint64, error) {
	v, err := Redis.Do("GET", key).Uint64()
	if err == redis.Nil {
		return 0, nil
	}
	return v, err
}

func GetFloat64(key string) (float64, error) {
	v, err := Redis.Do("GET", key).Float64()
	if err == redis.Nil {
		return 0.0, nil
	}
	return v, err
}

func TTL(key string) (int, error) {
	ttl, err := Redis.Do("TTL", key).Int()
	if err != nil {
		return -1, err
	}

	return ttl, nil
}

func Del(key string) error {
	err := Redis.Do("DEL", key).Err()
	if err != nil {
		loggers.Logs(fmt.Sprint("RedisDel Error! key:", key, "Details:", err.Error()))
	}
	return err
}

func RedisHGet(key, field string) (string, error) {
	value, err := Redis.Do("HGET", key, field).String()
	if err != nil {
		return "", nil
	}

	return value, nil
}

func HSet(key, field, value string) error {
	err := Redis.Do("HSET", key, field, value).Err()
	if err != nil {
		loggers.Logs(fmt.Sprint("RedisHSet Error!", key, "field:", field, "Details:", err.Error()))
	}
	return err
}

func RedisHDel(key, field string) error {
	err := Redis.Do("HDEL", key, field).Err()
	if err != nil {
		loggers.Logs(fmt.Sprint("RedisHDel Error!", key, "field:", field, "Details:", err.Error()))
	}
	return err
}

func ZAdd(key, member, score string) error {
	err := Redis.Do("ZADD", key, score, member).Err()
	if err != nil {
		loggers.Logs(fmt.Sprint("RedisZAdd Error!", key, "member:", member, "score:", score, "Details:", err.Error()))
	}
	return err
}

func RedisZRank(key, member string) (int, error) {
	rank, err := Redis.Do("ZRANK", key, member).Int()
	if err == redis.Nil {
		return -1, nil
	}

	if err != nil {
		loggers.Logs(fmt.Sprint("RedisZRank Error!", key, "member:", member, "Details:", err.Error()))
		return -1, nil
	}

	return rank, err
}

func ZRange(key string, start, stop int) (values []string, err error) {
	values, err = Redis.ZRange(key, int64(start), int64(stop)).Result()
	if err != nil {
		loggers.Logs(fmt.Sprint("RedisZRange Error!", key, "start:", start, "stop:", stop, "Details:", err.Error()))
		return
	}

	return
}

func ZRangeWithScores(key string, start, stop int) (values []redis.Z, err error) {
	values, err = Redis.ZRangeWithScores(key, int64(start), int64(stop)).Result()
	if err != nil {
		loggers.Logs(fmt.Sprint("RedisZRange Error!", key, "start:", start, "stop:", stop, "Details:", err.Error()))
		return
	}

	return
}

func ZRem(key, member string) error {
	err := Redis.Do("ZREM", key, member).Err()
	if err != nil {
		loggers.Logs(fmt.Sprint("RedisZRem Error!", key, "member:", member, "Details:", err.Error()))
	}
	return err
}

func RPUSH(key string, member string) (err error) {
	err = Redis.Do("RPUSH", key, member).Err()
	if err != nil {
		loggers.Logs(fmt.Sprint("RedisRPUSH Error!", key, member, "Details:", err.Error()))
		return
	}

	return
}

func LPOP(timeout time.Duration, keys ...string) (value []string, err error) {
	value, err = Redis.BLPop(timeout, keys...).Result()
	if err == redis.Nil {
		err = nil
		return
	}

	if err != nil {
		loggers.Logs(fmt.Sprint("BLPop Error!", keys, timeout, "Details:", err.Error()))
		return
	}
	return
}

func LLEN(key string) (value int64, err error) {
	value, err = Redis.LLen(key).Result()
	if err != nil {
		loggers.Logs(fmt.Sprint("RedisLLEN Error!", key, "Details:", err.Error()))
		return
	}

	return
}

func LRange(key string, start, stop int) (values []string, err error) {
	values, err = Redis.LRange(key, int64(start), int64(stop)).Result()
	if err != nil {
		loggers.Logs(fmt.Sprint("RedisLRange Error!", key, "start:", start, "stop:", stop, "Details:", err.Error()))
		return
	}

	return
}

func Keys(pattern string) (keys []string, err error) {
	keys, err = Redis.Keys(pattern).Result()
	if err != nil {
		loggers.Logs(fmt.Sprint("RedisKeys Error!", pattern, "Details:", err.Error()))
		return
	}

	return
}

// getKeys will take a certain prefix that the keys share and return a list of all the keys
func getKeys(prefix string) ([]string, error) {
	var allKeys []string
	var cursor uint64
	count := int64(10) // count specifies how many keys should be returned in every Scan call

	for {
		var keys []string

		keys, cursor, _ = Redis.Scan(cursor, prefix, count).Result()

		allKeys = append(allKeys, keys...)

		if cursor == 0 {
			break
		}

	}

	return allKeys, nil
}

func BatchDel(key ...string) error {
	err := Redis.Del(key...).Err()
	if err != nil {
		loggers.Logs(fmt.Sprint("RedisBatchDel Error! key:", key, "Details:", err.Error()))
	}
	return err
}

func Mset(pairs ...interface{}) error {
	err := Redis.MSet(pairs...).Err()
	if err != nil {
		loggers.Logs(fmt.Sprint("RedisMset Error! pairs:", pairs, "Details:", err.Error()))
	}
	return err
}
