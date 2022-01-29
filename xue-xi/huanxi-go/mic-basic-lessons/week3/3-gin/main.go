package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	productGroup := r.Group("/product")
	{
		productGroup.GET("/:id/:name", productDetail)
	}
	r.Run()

}

func productDetail(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "参数错误",
			"name": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  id,
			"name": name,
		})
	}
}
