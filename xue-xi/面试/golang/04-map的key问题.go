package main

import "fmt"

/**
 golang 中 map 的 key 必须是可比较的，再简单点是可以使用 == 运算符进行比较。
 很显然 slice，map，function 不可以，
 所以 数字、string、bool、array、channel、指针可以，以及 包含前面类型的 struct
 */


func main () {
	//可以
	 //var mstr map[string]string
	 //fmt.Println(mstr)
	//可以
	//var mbool map[bool]string
	//fmt.Println(mbool)
	//可以
	//var marr map[[1]int]string
	//fmt.Println(marr)


	 //不可以
	//var mslice map[[]int]string
	//fmt.Println(mslice)

	t1()
}


func t1() {
	var m1 map[string]string


	//m := make(map[string]string)

	fmt.Println("=====" ,m1["1"])
}