package internal

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

var RedisClient *redis.Client

func initRedis() {
	h := AppConf.RedisConfig.Host
	p := AppConf.RedisConfig.Port
	addr := fmt.Sprintf("%s:%d", h, p)
	fmt.Println(addr)
	fmt.Println(h)
	fmt.Println(p)

	RedisClient = redis.NewClient(&redis.Options{
		Addr: addr,
	})
	ping := RedisClient.Ping(context.Background())
	fmt.Println(ping)
	fmt.Println("redis初始化完成...")

	/*	h := config.RedisConfig.Host
		p := config.RedisConfig.Port
		fmt.Println(h)
		fmt.Println("my-redis")
		fmt.Println(p)*/

}
