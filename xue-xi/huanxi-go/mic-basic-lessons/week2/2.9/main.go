package main

import (
	"errors"
	"fmt"
)

func funcRecover() error {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("recover panic! v :%v \n", v)
		}
	}()
	return funcCook()
}

func funcCook() error {
	panic("停水\n")
	return errors.New("错误了")
}

//错误处理
func main() {
	//pain recover defer
	err := funcRecover()
	if err != nil {
		fmt.Printf("err v :%v", err)
	} else {
		fmt.Printf("err is nil")
	}

}
