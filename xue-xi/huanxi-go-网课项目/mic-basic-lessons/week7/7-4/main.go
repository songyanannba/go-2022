package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	v.SetConfigFile("week7/7-4/dev-config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(v.GetString("name"))

	//读取到结构体
	var dbConfig DBConfig
	err = v.Unmarshal(&dbConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(dbConfig.UserName)
	fmt.Println(dbConfig.Password)

}


type DBConfig struct {
	UserName string `mapstructure:"userName"`
	Password string `mapstructure:"password"`
}