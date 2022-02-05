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
	h := ViperConf.RedisConfig.Host
	p := ViperConf.RedisConfig.Port
	addr := fmt.Sprintf("%s:%d", h, p)
	fmt.Println(addr)

	RedisClient = redis.NewClient(&redis.Options{
		Addr: addr,
	})
	ping := RedisClient.Ping(context.Background())
	fmt.Println(ping)
	fmt.Println("redis初始化完成...")

}
