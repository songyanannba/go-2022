package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"mic-basic-lessons/week6/6-6/proto/pb"
	"net"
	"time"
)

type ToDo struct {
}

func (t *ToDo) DoWork(ctx context.Context, req *pb.TodoRequest) (*pb.TodoResponse, error) {
	fmt.Println("...start...")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("没有md")
	}
	for k, v := range md {
		fmt.Printf("key=%s: value=%s\n", k, v)
	}
	time.Sleep(time.Second * 2)
	fmt.Println(req.Todo + "已完成")
	return &pb.TodoResponse{
		Done: true,
	}, nil
}

func MyInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	now := time.Now()
	fmt.Println(now)
	handler(ctx, req)
	d := time.Now().Sub(now)
	fmt.Printf("执行时间%d", d.Milliseconds())
	return
}

func main() {

	serverInterceptor := grpc.UnaryInterceptor(MyInterceptor)
	server := grpc.NewServer(serverInterceptor)
	pb.RegisterToDoServiceServer(server, &ToDo{})

	listen, err := net.Listen("tcp", "0.0.0.0:9095")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}

}
