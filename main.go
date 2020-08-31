package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}


func main() {
	err := initClient()
	if err != nil {
		log.Println("连接失败")
		return
	}
	fmt.Println("连接成功")
}
