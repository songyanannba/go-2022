package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("hello", func(c *gin.Context) {
		lesson := c.PostForm("lesson")
		c.JSON(http.StatusOK, gin.H{
			"msg": "post 课程" + lesson,
		})
	})
}
