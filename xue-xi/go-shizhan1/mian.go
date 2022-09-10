package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)




func main() {
	fmt.Println("start service...")
	httpService()
	//GinService()
}

func httpService() {

	http.HandleFunc("/login", login)         //设置访问的路由
	fmt.Println("httpService succ...")
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


func GinService () {

	router := gin.Default()
	router.GET("/ping" , func(c *gin.Context) {
		c.JSON(200 , gin.H{
			"message" : "pong",
		})
	})
	router.Run(":8800")
}








func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, err := template.ParseFiles("forms/login.html")
		if err != nil {
			fmt.Println("template err" , err)
			return
		}
		log.Println(t.Execute(w, nil))

	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		_ = r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		if pwd := r.Form.Get("password"); pwd == "123456" { // 验证密码是否正确
			fmt.Fprintf(w, "欢迎登陆，Hello %s!", r.Form.Get("username")) //这个写入到w的是输出到客户端的
		} else {
			fmt.Fprintf(w, "密码错误，请重新输入!")
		}
	}
}