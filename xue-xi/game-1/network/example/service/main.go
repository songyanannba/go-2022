package main

import (
	"fmt"
	"game-1/network"
)

func main() {
	service := network.NewService(":8088", "tcp")
	service.Run()

	fmt.Println("service run ...")
	select {}
}
