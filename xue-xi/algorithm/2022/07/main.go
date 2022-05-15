package main

import (
	"fmt"
	"time"
)

type treeList struct {
	zuo *treeList
	you *treeList
	val string
}

var list []treeList
var zhan []*treeList

func add(v string) {
	fmt.Println("val is :", v)
}

func accessTree(tList *treeList) {
	if tList == nil {
		return
	}
	//time.Sleep(time.Second)
	//fmt.Printf("tlist is :%v \n" ,tList)
	//add(tList.val) //先
	accessTree(tList.zuo)
	add(tList.val) //中序遍历
	accessTree(tList.you)
	//add(tList.val) //后序遍历
}

func createTree() treeList {
	l5 := treeList{
		zuo: nil,
		you: nil,
		val: "5",
	}
	l6 := treeList{
		zuo: nil,
		you: nil,
		val: "6",
	}

	l7 := treeList{
		zuo: nil,
		you: nil,
		val: "7",
	}
	l4 := treeList{
		zuo: nil,
		you: nil,
		val: "4",
	}

	l2 := treeList{
		zuo: &l7,
		you: &l4,
		val: "2",
	}
	l3 := treeList{
		zuo: &l5,
		you: &l6,
		val: "3",
	}

	l1 := treeList{
		zuo: &l2,
		you: &l3,
		val: "1",
	}
	return l1
}

func digui(l1 treeList) {
	//fmt.Println(l1)
	//fmt.Printf("zhi %v" , l1)
	accessTree(&l1)
}

func xunhuandiedai(root *treeList) {

	for root != nil || len(zhan) > 0 {
		for root != nil {
			zhan = append(zhan, root)
			root = root.zuo
		}
		root = zhan[len(zhan)-1]
		zhan = zhan[:len(zhan)-1]
		add(root.val)
		time.Sleep(time.Second)
		root = root.you
	}

}

func main() {

	tree := createTree()
	//树遍历
	//digui(tree) //递归
	xunhuandiedai(&tree) //循环迭代

}
