package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/sss", func(c *gin.Context) {
		fmt.Println("...sss")
		lesson := c.PostForm("lesson")
		c.JSON(http.StatusOK, gin.H{
			"msg": "post 课程" + lesson,
		})
	})
	r.Run()
}
