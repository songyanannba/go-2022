package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CrossDomain(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSrF-Token,Authorization,Token,x-Token")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Expose-Methods", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}
