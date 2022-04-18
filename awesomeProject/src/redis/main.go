package main

import (
	"github.com/go-redis/redis"

)

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr: "redis://localhost:6379",
		Password: "",
		DB: 0,
	})
	_, err = redisdb.Ping().Result()
	return
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Println("redis error")
		return
	}
	fmt.Println("连接redis成功")
}