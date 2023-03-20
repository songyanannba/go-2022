package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Cook() {
	now := time.Now()
	fmt.Println("开始做饭")
	time.Sleep(time.Second * 3)
	fmt.Println("结束做饭")
	expired := time.Now().Sub(now)
	fmt.Println(expired)
}

func CustomMiddleWare(c *gin.Context) {
	now := time.Now()
	c.Next()
	expired := time.Now().Sub(now)
	fmt.Println(expired)
}

func Buy() {
	fmt.Println("..买菜")
	time.Sleep(time.Second * 2)
	fmt.Println("..买菜回来")
}

func Cook2() {
	fmt.Println("开始做饭")
	time.Sleep(time.Second * 3)
	fmt.Println("结束做饭")
}

func Eat() {
	fmt.Println("..吃饭")
	time.Sleep(time.Second * 1)
	fmt.Println("..吃完了")
}

func Wash() {
	fmt.Println("..洗碗")
	time.Sleep(time.Second * 1)
	fmt.Println("..洗完了")
}

func main() {
	//Cook()
	r := gin.Default()
	r.Use(CustomMiddleWare)
	r.GET("/", func(c *gin.Context) {
		Buy()
		Cook2()
		Eat()
		Wash()
	})
	r.Run()
}
