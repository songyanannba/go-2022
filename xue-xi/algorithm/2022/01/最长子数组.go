package a01

import "fmt"

/*给定一个长度为n的数组arr，返回arr的最长无重复元素子数组的长度，无重复指的是所有数字都不相同。
子数组是连续的，比如[1,3,5,7,9]的子数组有[1,3]，[3,5,7]等等，但是[1,3,7]不是子数组

输入：
[2,3,4,5]
返回值：
4
说明：
[2,3,4,5]是最长子数组
要求：空间复杂度 O(n)O(n)，时间复杂度 O(nlogn)O(nlogn)*/

func MaxLength1(arr []int) int {
	res := 0
	s := make([]int, 100000)
	j := 0
	for  i := 0 ; i < len(arr); i++ {
		s[arr[i]]++
		for j <= i && s[arr[i]] > 1 {
			s[arr[j]] --
			j++
		}
		if i-j + 1 > res {
			res = i - j + 1
		}
	}
	return  res
}

func maxLength2(arr []int) int {
	left ,flag ,res := 0 ,make([]int, 100000) , 0
	for i:=0 ; i < len(arr) ; i++ {
		flag[arr[i]]++
		for left <= i && flag[arr[i]] > 1 {
			flag[arr[left]] --
			left ++
		}
		if i - left + 1 >res {
			res = i -left +1
		}
	}
	return res
}


func main(){
	var arr1 = []int {1,3,4,5 ,5,6}
	length1 := MaxLength1(arr1)
	length2 := maxLength2(arr1)
	fmt.Println(length1)
	fmt.Println(length2)
}








