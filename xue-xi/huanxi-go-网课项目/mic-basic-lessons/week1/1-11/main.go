package main

import "fmt"

//切片 ptr len cap

func SoldOut(foods []string) {
	foods[1] = "以卖完"
	fmt.Println(foods)
}

func main () {
	arr := [...]string{"烧乳猪" ,"鲍鱼" ,"大盘鸡" , "烧鹅" ,"烤鸭" ,"大龙虾"}

	//切片 左闭区间 右开区间
	slice1 := arr[1:3]
	fmt.Println(slice1)
	slice2 := arr[:3]
	fmt.Println(slice2)

	//全要
	fmt.Println(arr[:])

	slice3 :=arr[:]
	//切片是引用传递
	SoldOut(slice3)
	fmt.Println(slice3)


	slice4 := arr[2:5]
	fmt.Println(slice4)

	slice5 := slice4[2:4]
	fmt.Println(slice5)

	slice8 := []string{}
	fmt.Println(slice8 ,len(slice8), cap(slice8))
	for i := 0 ; i< 10 ; i++ {
		slice8 = append(slice8 ,fmt.Sprintf("但茶饭%d" , i))
		fmt.Println(slice8,len(slice8), cap(slice8))
	}
	slice8 = append(slice8 ,"但茶饭ddd")
	fmt.Println(slice8,len(slice8), cap(slice8))

	s1 := make([]string , 8)
	fmt.Println(s1)

}
