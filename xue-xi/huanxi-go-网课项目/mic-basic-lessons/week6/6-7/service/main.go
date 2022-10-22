package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"mic-basic-lessons/week6/6-6/proto/pb"
	"net"
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
	fmt.Println(req.Todo + "已完成")
	return &pb.TodoResponse{
		Done: true,
	}, nil
}

func main() {

	/*	md1 := metadata.New(map[string]string{
			"name": "sdjkhfajk",
		})
		md2 := metadata.Pairs("name", "jjjjj")*/

	server := grpc.NewServer()
	pb.RegisterToDoServiceServer(server, &ToDo{})
	listen, err := net.Listen("tcp", "0.0.0.0:9094")
	if err != nil {
		panic(err)
	}

	err1 := server.Serve(listen)
	if err1 != nil {
		panic(err)
	}

}
