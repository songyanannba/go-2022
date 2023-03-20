package main

import (
	"context"
	"google.golang.org/grpc"
	"mic-basic-lessons/week5/5-9/ptoto/pb"
	"net"
)

type BookInfo struct {
}

func (b *BookInfo) Study(ctx context.Context, request *pb.BoolRequest) (*pb.BookResponse, error) {
	return &pb.BookResponse{
		Msg: "w y x x" + request.Name,
	}, nil
}

func main() {

	server := grpc.NewServer()
	pb.RegisterStudyServer(server, &BookInfo{})
	listen, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		panic(nil)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(nil)
	}
}
