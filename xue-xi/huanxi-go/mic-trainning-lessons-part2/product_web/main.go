package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mic-training-lessons-part2/internal"
	"mic-training-lessons-part2/internal/register"
	"mic-training-lessons-part2/product_web/hander"
	"mic-training-lessons-part2/util"
)

var (
	consulRegistry register.ConsulRegistry
	randomId       string
)

func init() {

	randomPort := util.GenRandomPort()
	if !internal.AppConf.Debug {
		internal.AppConf.ProductWebConfig.Port = randomPort
	}

	randomId := uuid.New().String()
	consulRegistry := register.NewConsulRegistry(internal.AppConf.ProductWebConfig.Host,
		internal.AppConf.ProductWebConfig.Port)

	consulRegistry.Register(internal.AppConf.ProductWebConfig.SrvName, randomId,
		internal.AppConf.ProductWebConfig.Port,
		internal.AppConf.ProductWebConfig.Tags)

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
		productGroup.POST("/add", hander.AddHandler)
		productGroup.POST("/Update", hander.UpdateHandler)
		productGroup.GET("/detail/:id", hander.DetailHandler)
		productGroup.GET("/delete/:id", hander.DeleteHandler)
	}

	r.GET("/health", hander.HealthHandler)
	r.Run(addr)
}
