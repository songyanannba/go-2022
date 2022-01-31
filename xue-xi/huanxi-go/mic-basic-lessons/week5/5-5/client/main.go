package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	c, _ := net.Dial("tcp", "localhost:9090")

	reply := ""

	cliend := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(c))

	cliend.Call("FoodService.SayName", "九转大肠", &reply)

	fmt.Println(reply)

}
