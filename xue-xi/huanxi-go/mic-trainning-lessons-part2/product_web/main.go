package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mic-training-lessons-part2/product_web/hander"

	"mic-training-lessons-part2/internal"
)

func init() {
	err := internal.Reg(
		internal.AppConf.ProductWebConfig.Host,
		internal.AppConf.ProductWebConfig.SrvName,
		internal.AppConf.ProductWebConfig.SrvName,
		internal.AppConf.ProductWebConfig.Port,
		internal.AppConf.ProductWebConfig.Tags)
	if err != nil {
		panic(err)
	}
	fmt.Println("internal.Reg...")
}

func main() {
	ip := internal.AppConf.ProductWebConfig.Host
	port := internal.AppConf.ProductWebConfig.Port

	addr := fmt.Sprintf("%s:%d", ip, port)

	fmt.Println(addr)
	r := gin.Default()
	productGroup := r.Group("/v1/product")
	{
		productGroup.GET("/list", hander.ProductListHandler)
	}

	r.GET("/health", hander.HealthHandler)
	r.Run(addr)
}
