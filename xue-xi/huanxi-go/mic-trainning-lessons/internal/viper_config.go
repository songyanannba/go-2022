package internal

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
)

var AppConf AppConfig
var NacosConf NacosConfig
//var ViperConf ViperConfig

//var fileName = "./dev-config.yaml"
//var fileName = "/Users/songyanan/GolandProjects/go-2022/xue-xi/huanxi-go/mic-trainning-lessons/dev-config.yaml"
var fileName = "/Users/songyanan/GolandProjects/go-2022/xue-xi/huanxi-go/mic-trainning-lessons/pro-config.yaml"

func initNacos() {
	v := viper.New()
	v.SetConfigFile(fileName)
	v.ReadInConfig()
	v.Unmarshal(&NacosConf)
	fmt.Println(NacosConf)
}

func initFromNacos() {
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: NacosConf.Host,
			Port:   NacosConf.Port,
		},
	}
	clientConfig := constant.ClientConfig{
		NamespaceId:         NacosConf.NameSpace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "nacos/log",
		CacheDir:            "nacos/cache",
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfig,
		"clientConfig":  clientConfig,
	})

	if err != nil {
		panic(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: NacosConf.DataId,
		Group:  NacosConf.Group,
	})
	//fmt.Println(content)
	json.Unmarshal([]byte(content), &AppConf)
	//fmt.Println(ViperConf)
}

//第二种结合nacos读取配置的方法

func init() {
	initNacos()
	initFromNacos()
	fmt.Println("初始化成功...")
	initRedis()
}














/*type ViperConfig struct {
	DBConfig         DBConfig         `mapstructure:"db"`
	RedisConfig      RedisConfig      `mapstructure:"redis"`
	ConsulConfig     ConsulConfig     `mapstructure:"consul"`
	AccountSrvConfig AccountSrvConfig `mapstructure:"account_srv"`
	AccountWebConfig AccountWebConfig `mapstructure:"product_web"`
	NacosConfig      NacosConfig      `mapstructure:"nacos"`
}*/

//第一种读取配置方法
func init1() {
	/*v := viper.New()
	v.SetConfigFile(fileName)
	v.ReadInConfig()
	err := v.Unmarshal(&ViperConf)*/
	/*if err != nil {
		panic(err)
	}*/
	/*fmt.Println(fileName)
	fmt.Println(ViperConf)
	fmt.Println(ViperConf.ConsulConfig.Host)
	fmt.Println(fileName)
	fmt.Println("哈哈哈")
	fmt.Println(ViperConf.AccountWebConfig.SrvName)
	fmt.Println(ViperConf.AccountWebConfig.Host)*/
}
