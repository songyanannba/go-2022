package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"mic-trainning-lessons-part4/cartorder_web/hander"
	"mic-trainning-lessons-part4/internal"
	"mic-trainning-lessons-part4/internal/register"
	"mic-trainning-lessons-part4/util"

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
		internal.AppConf.CartOrderWebConfig.Port = randomPort
	}

	randomId := uuid.New().String()
	consulRegistry := register.NewConsulRegistry(internal.AppConf.CartOrderWebConfig.Host,
		internal.AppConf.CartOrderWebConfig.Port)

	consulRegistry.Register(internal.AppConf.CartOrderSrvConfig.SrvName, randomId,
		internal.AppConf.CartOrderWebConfig.Port,
		internal.AppConf.CartOrderWebConfig.Tags)

	fmt.Println("internal.Reg...")
}

func main() {
	ip := internal.AppConf.CartOrderWebConfig.Host
	port := internal.AppConf.CartOrderWebConfig.Port

	addr := fmt.Sprintf("%s:%d", ip, port)

	fmt.Println(addr)
	r := gin.Default()
	productGroup := r.Group("/v1/cart")
	{
		productGroup.GET("/list", hander.ShopCartListHandler)
		productGroup.POST("/add", hander.AddHandler)
		productGroup.POST("/Update", hander.UpdateHandler)
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
