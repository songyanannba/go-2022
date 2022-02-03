package internal

import (
	"github.com/spf13/viper"
	"mic-basic-lessons/week7/7-6/conf"
)

var ServerConfig conf.ServerConfig

func init() {
	v := viper.New()
	v.SetConfigFile("week7/7-6/conf/dev-config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.Unmarshal(&ServerConfig)

}
