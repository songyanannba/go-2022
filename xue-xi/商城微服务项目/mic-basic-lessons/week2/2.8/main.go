package main

import (
	"bufio"
	"fmt"
	"os"
)

func Cook() {
	//defer 先进后出 栈 （子弹上膛）
	//为什么要defer 1-orm 防止内存溢出
	defer fmt.Println("开饭")
	defer fmt.Println("播放音乐")
	fmt.Println("买菜")
	fmt.Println("卖肉")
	//panic("停水了")
	fmt.Println("做饭")
	fmt.Println("sss")
}

func WriteMee(fileName string, foods []string) {
	curDir, _ := os.Getwd()
	fmt.Println(curDir)
	path := curDir + fileName
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)
	defer w.Flush()
	for _, item := range foods {
		fmt.Fprintln(w, item)
	}
}

func main() {
	//defer
	//Cook()
	foods := []string{"烧鸡", "烤鸭", "羊腿", "干炒牛肉"}
	p := "/xue-xi/商城微服务项目/mic-basic-lessons/week2/5-2-http.8/foods.txt"
	WriteMee(p, foods)
}
