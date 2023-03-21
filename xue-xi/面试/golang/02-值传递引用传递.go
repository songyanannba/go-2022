package main

import (
	"fmt"
)

/**
	golang中都是采用值传递，即拷贝传递，也就是深拷贝。没有引用传递。
	之所有有些看起来像是引用传递的场景，是因为Golang中存在着引用类型，
		如slice、map、channel、function、pointer这些天生就是指针的类型，在传递的时候是复制的地址。
		引用类型作为参数时，称为浅拷贝，形参改变，实参数跟随变化。因为传递的是地址，形参和实参都指向同一块地址
 */

type User struct {
	name string
	age int
}

func a(u *User) {
	u.name = "sss"
	u.age = 20
	//fmt.Println("aa == " , u)
}

func ss(a []string) {
	a[0] = "M"
	a = append(a , "c") //这里注意 不会影响外面的切片
	a = append(a , "d")
	fmt.Println("slice == " , a)
}

func main() {

	u := User{
		name: "syn",
		age: 18,
	}

	fmt.Println("11" , u)
	a(&u)
	fmt.Println("22" , u)

	slc := []string{
		"a","b",
	}

	fmt.Println("11" , slc)
	ss(slc)
	fmt.Println("33" , slc)
}
