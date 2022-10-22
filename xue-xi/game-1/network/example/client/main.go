package main

import (
	"fmt"
	"game-1/network"
)

func main() {

	client := network.NewClient(":8088")
	client.Run()

	fmt.Println("client run ...")
	select {}

}
