package main

import "fmt"

/**
 总结
	1 如果有返回值 或者返回值为引用类型 defer里面的值会影响返回值
	2 如果没有返回值 defer里面的值操作 不影响返回值
 */

func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i //或者直接写成return
}

func d() *int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return &i
}

func c() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i
}

func f() int{
	defer func() {fmt.Println("111")}()
	defer func() {fmt.Println("22")}()
	defer func() {fmt.Println("333")}()

	return 1
	//panic("cccc")
}

func main() {
	//fmt.Println("return b :", b())  //影响
	///fmt.Println("return d :", *d()) //影响
	//fmt.Println("return c :", c())  //不影响

	//fmt.Println( "123 = " , f())
}