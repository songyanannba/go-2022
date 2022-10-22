package main

import "fmt"

//完全二叉树 要么是满二叉树 ，要么从左到右节点存在
// 1 有右节点 没有左节点 不是完全二叉树
// 2 在上面情况不违规的情况下， 遇到的第一个节点不双全，后续都是叶子节点

type Head struct {
	Prev, Next *Head
	Value      int
}

func IsCbt(Node *Head) bool {
	if Node == nil {
		return false
	}

	//放入队列 然后 遍历每个节点
	//todo

	return true
}

func GetTree() *Head {
	node := &Head{
		Value: 1,
	}
	node2 := &Head{
		Value: 2,
	}

	node3 := &Head{
		Value: 3,
	}
	node4 := &Head{
		Value: 4,
	}
	node5 := &Head{
		Value: 5,
	}
	node.Prev = node2
	node.Next = node3

	node2.Prev = node4
	node2.Next = node5
	return node
}

func main() {

	//构造树

	tree := GetTree()
	fmt.Println(tree)

	//逻辑处理
	// 1 有右节点 没有左节点 不是完全二叉树
	// 2 在上面情况不违规的情况下， 遇到的第一个节点不双全，后续都是叶子节点
	IsCbt(tree)
}
