package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"mag": "go-面向加薪学习",
	})
}

func main1() {
	r := gin.Default()
	//r := gin.New()

	r.GET("/", hello)
	r.Run()
}
