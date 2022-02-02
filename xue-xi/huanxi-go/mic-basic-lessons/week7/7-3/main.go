package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	r := gin.Default()
	/*pro, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	pro.Sugar()*/
	pro, _ := zap.NewProduction()
	zap.ReplaceGlobals(pro)
	port := 9090
	log := zap.S()
	defer log.Sync()
	log.Infof("服务端启动端口 %d \n", port)

	r.GET("/", func(c *gin.Context) {
		log.Info("sssddd...")
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})

	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Panic("无服务失败")
	}

}
