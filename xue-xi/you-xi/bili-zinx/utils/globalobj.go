package utils

import (
	"bili-zinx/zinx/ziface"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/**
	存储 框架全局参数 供其他模块使用
*/

type GlobalObj struct {
	TcpServer ziface.IService
	Host      string
	Port int
	Name string

	Version string
	MaxConn int
	MaxPackageSize uint32 //当前框架最大数据包

}

//定义全局对象
var GlobalObject *GlobalObj

func (g *GlobalObj)Reload () {
	dir, _ := os.Getwd()
	fmt.Println("dir==",dir)
	date, err := ioutil.ReadFile(dir+"/conf/zinx.json")
	if err != nil {
		panic(err)
	}
	//将文件数据 解析到globalobj里面

	err = json.Unmarshal(date, GlobalObject)
	if err != nil {
		fmt.Println("解析配置文件报错")
		panic(err)
	}
}


//初始化 当前对象
func init() {
	//默认值
	GlobalObject = &GlobalObj{
		Host:           "0.0.0.0",
		Port:           8999,
		Name:           "ZinxServerApp",
		Version:        "V0.4",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	//尝试从配置文件家在数据
	GlobalObject.Reload()
}
