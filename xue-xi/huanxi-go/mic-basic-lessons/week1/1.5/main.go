package main

import "fmt"


const (
	Sunday  = iota
	Monday
	Tuesday

	author = "syn"
	book = "Go"
)

//常量 枚举
func main () {

	const  version  = 1.1
	const appName = "sb"

	fmt.Println(version , appName ,author ,book)
	fmt.Println(Sunday , Monday ,Tuesday)

}
