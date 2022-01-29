package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"week3/1-gin1/model"
)

func main() {
	r := gin.Default()
	//账户
	accountGroup := r.Group("/account")
	{
		accountGroup.GET("/list", accountList)
		accountGroup.POST("/add", accountAdd)
	}
	//r.GET()

	//产品
	productGroup := r.Group("/product")
	{
		productGroup.GET("/productList", productList)
	}
	//订单
	r.Run()

}

func productList(c *gin.Context) {

}

func accountList(c *gin.Context) {
	var accountList []model.Account
	a1 := model.Account{
		No:   1,
		Name: "老王",
	}
	accountList = append(accountList, a1)
	a2 := model.Account{
		No:   2,
		Name: "老张",
	}
	accountList = append(accountList, a2)
	c.JSON(http.StatusOK, gin.H{
		"accountList": accountList,
	})
}

func accountAdd(c *gin.Context) {

}
