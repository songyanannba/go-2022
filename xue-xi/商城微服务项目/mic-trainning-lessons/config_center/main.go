package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"mic-trainning-lessons/internal"
)

//读取 nacos 的配置数据 ｜ 环境隔离
//https://github.com/nacos-group/nacos-sdk-go

func main() {

	nacosConfig := internal.NacosConf
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: nacosConfig.Host,
			Port:   nacosConfig.Port,
		},
	}

	clientConfig := constant.ClientConfig{
		NamespaceId:         nacosConfig.NameSpace,
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
		DataId: "account_json",
		Group:  "pro",
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}
