package main

import "fmt"

//变量命名规范
func one() {
	//1 var 变量名 变量类型
	var price int32 = 68
	fmt.Println(price)

	var name = "宫保鸡丁"
	fmt.Println(name)

	//2 定义多个变量
	var price1 ,weight = 68 ,1
	fmt.Println(price1 , weight)

	//省略类型 :=
	price3 := 66
	weight1 := 11
	fmt.Printf("peice3: %d ; weight : %d \n" , price3, weight1)

	//var 定义多个
	var (
		a int = 1
		b string = "bb"
	)
	fmt.Println(a , b)
}

//函数外面
var book = "GO语言"
var (
	aa int = 1111
	bb string = "BB"
	lesson1 = "0到GO架构师"
)

func main() {
	one()//变量命名规范

	fmt.Println(book)
	fmt.Println(aa,bb ,lesson1)
}
