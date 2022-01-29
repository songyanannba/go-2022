package main

import (
	"fmt"
	"sync"
)

func main() {
	s := []string{
		"回归肉",
		"ku",
		"快结束的功夫",
		"是的感觉",
	}

	var wg sync.WaitGroup

	for _, item := range s {
		wg.Add(1)
		go SayFoodName(item, &wg)
	}

	wg.Wait()
	fmt.Println("菜已经上起...")
}

func SayFoodName(name string, wg *sync.WaitGroup) {
	fmt.Println("您点的菜" + name)
	wg.Done()
}
