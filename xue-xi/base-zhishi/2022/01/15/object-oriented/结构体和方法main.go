package tree

import "fmt"

type TreeNode struct {
	Value int
	Left ,Right *TreeNode
}

//自定义构造函数 （工厂）
func creatNodes(value int) *TreeNode {
	//栈上 还是 堆上 根据运行环境的编辑器
	return &TreeNode{Value: value}
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
	root.Left = new(TreeNode)
	root.Right = new(TreeNode)
	root.Left.SetValue(200)

	root.Left.Right = new(TreeNode)
	root.Left.Right.SetValue(500)

	root.Right.SetValue(300)
	fmt.Println(root.Left)
	fmt.Println(root.Right)
	//root.left.right.setValue(400)
	//root.right.setValue(500)
	//root.right.right.setValue(300)
	root.Traverse()
}


func (node *TreeNode) Print() {
	if node == nil {
		fmt.Println("print value to nil ..")
		return
	}
	fmt.Println(node.Value)
}

func (node *TreeNode) SetValue (value int) {
	if node == nil {
		fmt.Println("setValue to nil ..")
		return
	}
	node.Value = value
}
//先序遍历
func (node *TreeNode) Traverse() {
	if node == nil {
		//fmt.Println("traverse")
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}


//xx1===
func Xx1 () {
	fmt.Println("start...")
	var root TreeNode
	fmt.Println(root)

	root = TreeNode{Value: 3}
	fmt.Println(root)

	root.Left = &TreeNode{}
	root.Right = &TreeNode{5,nil ,nil}
	fmt.Println(root)
	//切片
	nodes := []TreeNode {
		{Value: 6},
		{},
		{8, nil ,nil},
	}
	fmt.Println(nodes)
	creatNodes(2)
}