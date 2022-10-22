package main

import (
	"fmt"
	"sync"
	"time"
)

type Goods struct {
	v map[string]int
	m sync.Mutex
}

func (g *Goods) Inc(key string, num int) {
	g.m.Lock()
	defer g.m.Unlock()
	fmt.Printf("%d 库存数量增加 ，加锁\n", num)
	g.v[key] += num
	fmt.Printf("%d 库存数量增加完毕，解锁\n", num)
}

func (g *Goods) Value(key string) int {
	g.m.Lock()
	defer g.m.Unlock()
	fmt.Println("库存值上锁")
	return g.v[key]
}

func main() {
	mutex := sync.Mutex{}
	g := Goods{
		v: make(map[string]int),
		m: mutex,
	}
	for i := 0; i < 10; i++ {
		go g.Inc("榴莲", 2)
	}
	time.Sleep(time.Second * 1)
	fmt.Println(g.Value("榴莲"))

}
