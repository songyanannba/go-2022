package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"mic-training-lessons-part2/internal"
	"mic-training-lessons-part2/internal/register"
	"mic-training-lessons-part2/product_web/hander"
	"mic-training-lessons-part2/util"
	"os"
	"os/signal"
	"syscall"
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

	go func() {
		err := r.Run(addr)
		if err != nil {
			zap.S().Panic(addr + "启动失败" + err.Error())
		} else {
			zap.S().Info(addr + "启动成功")
		}
	}()
	q := make(chan os.Signal)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	<-q

	err := consulRegistry.DeRegister(randomId)

	if err != nil {
		zap.S().Panic("注销失败" + randomId + ":" + err.Error())
	} else {
		zap.S().Info("注销成功" + randomId)
	}
	//r.Run(addr)
}
