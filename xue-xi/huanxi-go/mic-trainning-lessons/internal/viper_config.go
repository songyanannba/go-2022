package internal

import (
	"fmt"
	"github.com/spf13/viper"
)

type ViperConfig struct {
	RedisConfig RedisConfig `mapstructure:"redis"`
}

var ViperConf ViperConfig

func init() {
	v := viper.New()
	configName := "dev-config.yaml"
	v.SetConfigFile(configName)
	v.ReadInConfig()
	err := v.Unmarshal(&ViperConf)
	if err != nil {
		panic(err)
	}
	fmt.Println(ViperConf)
	fmt.Println("初始化成功...")
	initRedis()
}
