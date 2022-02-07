package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"mic-trainning-lessons/account_web/hander"
	"mic-trainning-lessons/internal"
)

func init() {
	err := internal.Reg(
		internal.ViperConf.AccountWebConfig.Host,
		internal.ViperConf.AccountWebConfig.SrvName,
		internal.ViperConf.AccountWebConfig.SrvName,
		internal.ViperConf.AccountWebConfig.Port,
		internal.ViperConf.AccountWebConfig.Tags,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("internal.Reg...")
}

func main() {
	ip := flag.String("ip", "192.168.1.4", "输入Ip")
	port := flag.Int("port", 8081, "输入端口")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *ip, *port)

	r := gin.Default()

	accountGroup := r.Group("/v1/account")
	{
		accountGroup.GET("/list", hander.AccountListHandler)
		accountGroup.POST("/login", hander.LoginByPasswordHandler)
		accountGroup.GET("/captcha", hander.CaptchaHandler)
	}

	r.GET("/health", hander.HealthHandler)
	r.Run(addr)
}
