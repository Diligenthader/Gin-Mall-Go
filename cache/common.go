package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
	"strconv"
)

var (
	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("Redis config err :", err)
	}
	LoadRedisData(file)
	Redis()
}

// LoadRedisData Redis的初始化配置
func LoadRedisData(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

// Redis 连接
func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		//Password
		DB: int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}