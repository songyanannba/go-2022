package main

import "fmt"

func byValue(price int) {
	price += 20
}

func byPre (price *int) {
	*price += 20
}

func main()  {
	//指针
	var price  int  = 66
	var ptr *int = &price
	fmt.Println(ptr)
	*ptr = 88
	fmt.Println(price)

	//值传递
	byValue(price)
	fmt.Println(price)
	//指针传递
	byPre(&price)
	fmt.Println(price)

}
