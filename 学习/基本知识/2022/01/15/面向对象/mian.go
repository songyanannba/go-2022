package main

import "fmt"

type treeNode struct {
	value int
	left ,right *treeNode
}

func main() {
	fmt.Println("start...")
	var root treeNode
	fmt.Println(root)

	root = treeNode{value: 3}
	fmt.Println(root)

	root.left = &treeNode{}
	root.right = &treeNode{5,nil ,nil}

	fmt.Println(root)
}
