package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"mic-basic-lessons/week6/6-6/proto/pb"
	"time"
)



func main() {

	cliInter := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		now := time.Now()
		err := invoker(ctx , method ,req, reply , cc , opts...)
		d := time.Now().Sub(now)
		fmt.Printf("客户端 执行时间...%d\n",d.Milliseconds())
		return err
	}

	interceptor := grpc.WithUnaryInterceptor(cliInter)
	conn, err := grpc.Dial("127.0.0.1:9095", grpc.WithInsecure(),interceptor)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewToDoServiceClient(conn)
	md2 := metadata.Pairs("name", "jjjjj")
	ctx := metadata.NewOutgoingContext(context.Background(), md2)
	work, err := client.DoWork(ctx, &pb.TodoRequest{
		Todo: "我要学习",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(work.Done)
	fmt.Println("...end...")

}
