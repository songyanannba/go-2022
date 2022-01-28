package main

import "fmt"

//结构体和方法

type Food struct {
	No   int32
	Name string
}

func NewFood(no int32, name string) *Food {
	return &Food{
		No:   no,
		Name: name,
	}
}

//方法

func (f *Food) Show() {
	fmt.Println(f.No, f.Name)
}

func (f *Food) SetName(name string)  {
	f.Name = name
}

func main() {

	var f Food
	f = Food{
		No:   1,
		Name: "海参",
	}
	fmt.Println(f)

	f1 := Food{
		Name: "川菜",
		No:   2,
	}
	fmt.Println(f1)

	f2 := NewFood(3, "回锅肉")
	fmt.Println(f2)
	f2.Show()

	f2.SetName("肉回国")
	f2.Show()

}
