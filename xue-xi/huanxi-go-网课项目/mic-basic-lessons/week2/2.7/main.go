package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
)

type CheckOut func(int, int) int

func GetTotal(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

//函数别名
type GenerateRandom func() int

func RandomSum() GenerateRandom {
	a, b := rand.Intn(10), rand.Intn(20)
	return func() int {
		a, b = b, a+b
		return a
	}
}

func (g GenerateRandom) Read(p []byte) (n int, err error) {
	next := g()
	if next > 23 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("rrr %d\n", next)
	return strings.NewReader(s).Read(p)
}

func PrintRes(reader io.Reader) {
	sacnner := bufio.NewScanner(reader)
	for sacnner.Scan() {
		fmt.Println("sacnner.Text()" + sacnner.Text())
		//fmt.Println("sacnner.Text()")
	}
}

func main() {

	//匿名函数
	/*show2 := func() {
		fmt.Println("show222")
	}
	show2()
	//fmt.Println(show2)
	var checkout CheckOut = func(x, y int) int {
		return x + y
	}
	fmt.Println(checkout(11,22))*/
	/*	t := GetTotal(68)
		fmt.Println(t(1-orm))
		sum := GetTotal(100)
		fmt.Println(sum(t(1-orm)))*/

	r := RandomSum()
	PrintRes(r)

}
