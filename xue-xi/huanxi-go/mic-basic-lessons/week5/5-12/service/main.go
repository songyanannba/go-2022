package main

import (
	"fmt"
	"google.golang.org/grpc"
	"mic-basic-lessons/week5/5-12/proto/pb"
	"net"
	"sync"
	"time"
)

type FoodInfo struct {
}

func (f *FoodInfo) SayName(request *pb.FoodStreamRequest, server pb.FoodService_SayNameServer) error {
	fmt.Println("SayName 已请求")
	server.Send(&pb.FoodStreamResponse{
		Msg: "您点的菜是：" + request.Name,
	})
	return nil
}

func (f *FoodInfo) PostName(server pb.FoodService_PostNameServer) error {
	fmt.Println("PostName 已请求")
	for {
		recv, err := server.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(recv)
	}
	return nil
}

func (f *FoodInfo) FullStream(server pb.FoodService_FullStreamServer) error {
	fmt.Println("FullStream 已请求")

	//server.Send()
	//server.Recv()
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan string, 5)

	go func() {
		defer wg.Done()
		for {
			item, err := server.Recv()
			if err != nil {
				fmt.Println(err)
			}
			c <- item.Name
			fmt.Println("已下单" + item.Name)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			foodName := <-c
			err := server.Send(&pb.FoodStreamResponse{Msg: "菜" + foodName + "做好了"})
			if err != nil {
				fmt.Println(err.Error())
			}
			time.Sleep(time.Second * 1)
		}
	}()

	wg.Wait()
	return nil
}

func main() {
	//监听端口
	listen, err := net.Listen("tcp", ":9091")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	pb.RegisterFoodServiceServer(server, &FoodInfo{})

	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
	fmt.Println("...end")
}
