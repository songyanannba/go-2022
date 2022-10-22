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

var fileName = "/Users/songyanan/GolandProjects/go-2022/xue-xi/huanxi-go-网课项目/mic-trainning-lessons-part3/dev-config.yaml"

func initNacos() {
	v := viper.New()
	v.SetConfigFile(fileName)
	v.ReadInConfig()
	v.Unmarshal(&NacosConf)
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
	json.Unmarshal([]byte(content), &AppConf)
}

//第二种结合nacos读取配置的方法

func init() {
	initNacos()
	initFromNacos()
	fmt.Println("初始化成功...")
	initRedis()
	InitDB()
}