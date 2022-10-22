package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type FoodService struct {
}

func (f *FoodService) SayName(request string, resp *string) error {
	*resp = "菜.." + request
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		return
	}
	err = rpc.RegisterName("FoodService", &FoodService{})
	if err != nil {
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			return
		}
		//rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
		fmt.Println("112233")
	}
}
