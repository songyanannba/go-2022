package main

import "fmt"

func  main() {
	//浮点型
	var price float64 = 99.90
	fmt.Println(price)

	//string
	var name = "哈哈哈"
	fmt.Println(name)

	//布尔
	var b = false
	fmt.Println(b)

	//byte
	fmt.Println([]byte(name))

	//rune == char ;  int32别名

	name ,price ,num  := "红烧肉" , 66 ,1
	var total = 0
	discount := 0.7
	total =  int(float64(price) * float64(num) * discount)
	fmt.Println(name ,total)
}
