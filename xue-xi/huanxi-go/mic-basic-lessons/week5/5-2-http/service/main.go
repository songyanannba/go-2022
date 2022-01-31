package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		lesson := c.Query("lesson")
		c.JSON(http.StatusOK, gin.H{
			"msg": "我要学习" + lesson,
		})
	})

	r.Run()
}
