package main

import (
	"fmt"
	"time"
)

func ShowBook() {
	fmt.Println("book kjhdsgfakjsdkfhjak")
}

func main() {

	//go ShowBook()
	/*for i := 0 ; i < 10 ; i++ {
		fmt.Println(fmt.Printf("i am %d \n" , i))
	}*/

	for i := 0; i < 10; i++ {
		go func(j int) {
			fmt.Println(fmt.Printf("i am %d \n", j))
		}(i)
	}

	time.Sleep(time.Second * 1)

}
