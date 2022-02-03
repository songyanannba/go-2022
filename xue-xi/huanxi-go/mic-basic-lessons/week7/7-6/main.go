package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mic-basic-lessons/week7/7-6/internal"
	"net/http"
)

func main() {
	r := gin.Default()

	h := internal.ServerConfig.GinConfig.Host
	p := internal.ServerConfig.GinConfig.Port

	addr := fmt.Sprintf("%s:%d", h, p)
	/*fmt.Println(internal.ServerConfig.ServerName)
	fmt.Println(h)
	fmt.Println(p)
	fmt.Println(addr)*/
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": addr,
		})
	})
	r.Run(addr)
}
