package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mic-trainning-lessons/account_srv/biz"
	"mic-trainning-lessons/account_srv/proto/pb"
	"mic-trainning-lessons/internal"
	"net"
)

func init() {
	internal.InitDB()
}

func main() {
	ip := flag.String("ip", "127.0.0.1", "输入IP")
	port := flag.Int("port", 9095, "输入端口")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *ip, *port)
	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server, &biz.AccountServer{})
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		zap.S().Error("account_srv异常..." + err.Error())
		panic(err)
	}
	fmt.Println("suc...")
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
