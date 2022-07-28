package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	loggin "github.com/sirupsen/logrus"
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

func RedisInit() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("Redis配置文件加载失败.", err)
	}
	LoadRedis(file)
	Redis()
}

func LoadRedis(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		DB:       int(db),
		Password: RedisPw,
	})
	_, err := client.Ping().Result()
	if err != nil {
		loggin.Info(err)
		panic(err)
	}
	RedisClient = client
}
