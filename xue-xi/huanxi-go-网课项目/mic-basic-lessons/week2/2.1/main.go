package main

import "fmt"

type Where interface {
	WhereDinner()
}

type GoToHappy interface {
	GoToDinner(dest string)
}

func Happy(w Where, h GoToHappy, dest string) {
	w.WhereDinner()
	h.GoToDinner(dest)
}

type Body struct {
	Name string
}

func (b *Body) WhereDinner() {
	fmt.Println(b.Name + "： girl 去哪里吃")
}

type Girl struct {
	Name string
}

func (g *Girl) GoToDinner(dest string) {
	fmt.Println(g.Name + "： body 去" + dest)
}

func main() {
	b := Body{"小明"}
	g := Girl{"小蓝"}
	Happy(&b , &g,"野炊")
}
