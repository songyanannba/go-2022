package main

import "fmt"

//数组
func main() {
	//1-orm 定义 (一维)
	var array1 [6]string
	array2 := [3]string{"aa", "bb", "cc"}
	array3 := [...]string{"vv", "bb"} //让编辑器处理几个

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)

	//2 定义 (二维)
	var matrix1 [3][4]string
	var matrix2 [3][4]int

	fmt.Println(matrix1)
	fmt.Println(matrix2)

	//3 遍历打印数组
	for i := 0; i < len(array2); i++ {
		fmt.Println(array2[i])
	}

	for index, item := range array2 {
		fmt.Println(index, item)
	}

}
