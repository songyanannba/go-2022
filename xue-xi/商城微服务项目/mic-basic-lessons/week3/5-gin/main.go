package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

func main() {

	r := gin.Default()
	productGroup := r.Group("/product")
	{
		productGroup.GET("/detail", getDetailHandler)
		productGroup.POST("/detail", detailHandler)
		productGroup.POST("/add", addHandler)
	}
	r.Run()

}

func detailHandler(c *gin.Context) {
	fmt.Println("post")
	id := c.DefaultPostForm("id", "0")
	name := c.DefaultQuery("name", "postDetailHandler")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}

func getDetailHandler(c *gin.Context) {
	fmt.Println("get")
	id := c.DefaultQuery("id", "0")
	name := c.DefaultQuery("name", "defaultName")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}

func addHandler(c *gin.Context) {
	id := rand.Intn(1000)
	name := c.DefaultPostForm("name", "defaultPostName")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}
