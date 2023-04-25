package main

import "fmt"

func main() {
	var a []string
	var b  = []string{}

	//fmt.Printf("%T,%V \n", a,a)
	fmt.Printf("%T %T\n", a,b)
	fmt.Printf("%#v %#v\n", a,b)

	fmt.Println(len(b))
	b = append(b ,"aaa")
	fmt.Println(b)
	fmt.Println(len(a))
	a = append(a ,"bbb")
	fmt.Println(b)
}
