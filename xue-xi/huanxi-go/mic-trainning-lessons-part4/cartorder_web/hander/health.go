package hander

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//健康检查
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}