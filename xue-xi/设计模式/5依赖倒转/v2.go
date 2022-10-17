package main


import "fmt"

//---抽象层

type Car interface {
	Run1()
}

type Driver interface {
	Drive(car Car)
}


//--实现层
type BenZ struct {

}

func (b *BenZ)Run1() {
	fmt.Println("奔驰 runing")
}

type BaoM struct {

}
func (bm *BaoM)Run1() {
	fmt.Println("宝马 runing")
}


type Zhang33 struct {

}

func (z *Zhang33)Drive(car Car)  {
	fmt.Println("张三 开")
	car.Run1()
}

type Li44 struct {

}

func (l *Li44)Drive(car Car)  {
	fmt.Println("李四 开")
	car.Run1()
}


//--业务逻辑层

//优化-依赖倒转
func V2() {
	var benz Car
	benz = new(BenZ)

	var z3 Zhang33
	z3.Drive(benz) //依赖于抽象层
}