package internal

import "github.com/spf13/viper"

type Config struct {
	RedisConfig RedisConfig `mapstructure:"redis"`
}
var config Config

func init() {
	v := viper.New()
	//v.SetConfigFile("mic-trainning-lessons/dev-config.yaml")
	v.SetConfigFile("/Users/songyanan/GolandProjects/go-2022/xue-xi/huanxi-go-网课项目/mic-trainning-lessons/dev-config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.Unmarshal(&config)

}