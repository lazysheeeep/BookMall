package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gopkg.in/ini.v1"
	"strconv"
	"time"
)

var (
	ctx         = context.Background()
	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func Init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("redis config error")
	}
	LoadRedisData(file)
	Redis()
}

func LoadRedisData(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

func Redis() {

	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		//redis没有设置初始密码
		Password: RedisPw,
		DB:       int(db),
		//超时设置
		DialTimeout:  5 * time.Second, //连接建立超时时间，默认5s
		ReadTimeout:  3 * time.Second, //读超时，默认3s，设置-1则取消读超时
		WriteTimeout: 3 * time.Second, //写超时，默认3s，设置-1
		PoolTimeout:  4 * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒

	})
	_, err := client.Ping(ctx).Result() //这里没有上下文context
	if err != nil {
		panic("redis连接失败")
	}
	//所有不能检测是否连接成功
	RedisClient = client
}
