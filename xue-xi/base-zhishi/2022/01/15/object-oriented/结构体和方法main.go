package tree

import "fmt"

type TreeNode struct {
	value int
	left ,right *TreeNode
}

//自定义构造函数 （工厂）
func creatNodes(value int) *TreeNode {
	//栈上 还是 堆上 根据运行环境的编辑器
	return &TreeNode{value: value}
}


/*func main() {
	//初始化和基本操作
	//xx1()
	//结构体方法
	xx2()

}*/


//xx2===
func Xx2() {
	//如果要改变值，需要时指针接收者
	fmt.Println("xx2")
	var root TreeNode
	root.SetValue(100)
	fmt.Println(root)
	root.left = new(TreeNode)
	root.right = new(TreeNode)
	root.left.SetValue(200)

	root.left.right = new(TreeNode)
	root.left.right.SetValue(500)

	root.right.SetValue(300)
	fmt.Println(root.left)
	fmt.Println(root.right)
	//root.left.right.setValue(400)
	//root.right.setValue(500)
	//root.right.right.setValue(300)
	root.Traverse()
}

func (node *TreeNode) print() {
	if node == nil {
		fmt.Println("print value to nil ..")
		return
	}
	fmt.Println(node.value)
}

func (node *TreeNode) SetValue (value int) {
	if node == nil {
		fmt.Println("setValue to nil ..")
		return
	}
	node.value = value
}

func (node *TreeNode) Traverse() {
	if node == nil {
		fmt.Println("traverse")
		return
	}
	node.left.Traverse()
	node.print()
	node.right.Traverse()
}


//xx1===
func Xx1 () {
	fmt.Println("start...")
	var root TreeNode
	fmt.Println(root)

	root = TreeNode{value: 3}
	fmt.Println(root)

	root.left = &TreeNode{}
	root.right = &TreeNode{5,nil ,nil}
	fmt.Println(root)
	//切片
	nodes := []TreeNode {
		{value: 6},
		{},
		{8, nil ,nil},
	}
	fmt.Println(nodes)
	creatNodes(2)
}