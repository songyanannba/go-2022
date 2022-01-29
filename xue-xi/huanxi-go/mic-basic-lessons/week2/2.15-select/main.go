package main

import (
	"fmt"
	"time"
)

func main() {

	/*select {
	case <- ch1:
	case data:= <- ch1:
	case ch3 <- 11:
	default data:= <- ch1:
	}*/

	start := time.Now()
	fmt.Println(start)

	ch1 := make(chan interface{})
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func() {
		time.Sleep(4 * time.Second)
		close(ch1)
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- "0-架构师"
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch3 <- "222讲"
		//close(ch3)
	}()

	fmt.Println("等待处理....")

	select {
	case <-ch1:
		fmt.Printf("未zuse... %v \n", time.Since(start))
	case c2 := <-ch2:
		fmt.Println("c2" + c2)
	case c3 := <-ch3:
		fmt.Println("c3" + c3)
	/*default:
		fmt.Println("select end")*/
	}

}
