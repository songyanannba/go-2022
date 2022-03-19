package main

import (
	"fmt"
)

func cf1(arr [10]int) {
	fmt.Println(arr)
	var newArr [10]int
	for i := 0; i < len(arr); i++ {
		newArr[arr[i]]++
		if newArr[arr[i]] >= 2 {
			fmt.Println("true")
			break
		}
		fmt.Println("false")
	}

}

func cf2(arr [10]int) {
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] == arr[arr[i]] {
			fmt.Println("true")
			break
		}
		temp := arr[arr[i]]
		arr[arr[i]] = arr[i] //回到本来应该取得位置
		arr[i] = temp
		fmt.Println(arr)
		fmt.Println("false")
		i--
	}

}

//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，
//但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
//输入： [2, 3, 1, 0, 2, 5, 3]
//输出：2 或 3
func main() {

	arr := [10]int{2, 9, 5, 7, 8, 1, 2, 3, 6, 3}
	cf1(arr)
	cf2(arr)

}
