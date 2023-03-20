package main

import (
	"fmt"
)

func GivenFood() chan string {
	ch := make(chan string)
	go func() {
		ch <- "回锅肉"
		ch <- "大肠"
		ch <- "dae"
		close(ch)
	}()
	return ch
}

func main() {
	/*ch := make(chan string)
	ch1 := make(chan string ,6)
	ch <- "回锅肉"
	<-ch1
	close(ch)
	<- ch*/

	ch := make(chan string)
	ch = GivenFood()
	/*for {
		if name, ok := <-ch; ok {
			fmt.Println(name)
		} else {
			break
		}
	}*/

	for data := range ch {
		fmt.Println(data)
	}
}
