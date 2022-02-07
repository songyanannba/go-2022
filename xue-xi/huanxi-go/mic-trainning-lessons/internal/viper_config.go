package internal

import (
	"fmt"
	"github.com/spf13/viper"
)

var ViperConf ViperConfig
//var fileName = "./dev-config.yaml"
var fileName = "/Users/songyanan/GolandProjects/go-2022/xue-xi/huanxi-go/mic-trainning-lessons/dev-config.yaml"


func init() {
	v := viper.New()
	v.SetConfigFile(fileName)
	v.ReadInConfig()
	err := v.Unmarshal(&ViperConf)
	if err != nil {
		panic(err)
	}
	/*fmt.Println(fileName)
	fmt.Println(ViperConf)
	fmt.Println(ViperConf.ConsulConfig.Host)
	fmt.Println(fileName)
	fmt.Println("哈哈哈")
	fmt.Println(ViperConf.AccountWebConfig.SrvName)
	fmt.Println(ViperConf.AccountWebConfig.Host)*/
	fmt.Println("初始化成功...")
	initRedis()
}

type ViperConfig struct {
	DBConfig DBConfig `mapstructure:"db"`
	RedisConfig RedisConfig `mapstructure:"redis"`
	ConsulConfig ConsulConfig `mapstructure:"consul"`
	AccountSrvConfig AccountSrvConfig `mapstructure:"account_srv"`
	AccountWebConfig AccountWebConfig `mapstructure:"account_web"`
}

