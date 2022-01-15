package main

import (
	"fmt"
	tree "go-2022/xue-xi/base-zhishi/2022/01/15/object-oriented"
)


type myTreeNode struct {
	node *tree.TreeNode
}

//后序遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil ||  myNode.node == nil{
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}


func xx3() {
	var root tree.TreeNode

	root.SetValue(3)
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{}
	root.Left.SetValue(200)
	root.Left.Right = &tree.TreeNode{}
	root.Left.Right.SetValue(500)
	root.Right.SetValue(300)

	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}

//面向对象
func main() {
	//初始化和基本操作
	//xx1()
	//结构体方法
	tree.Xx2()
	//后序遍历
	xx3()

}