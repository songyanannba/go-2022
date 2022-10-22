package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"mic-basic-lessons/week6/6-6/proto/pb"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:9094", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewToDoServiceClient(conn)

	/*	md1 := metadata.New(map[string]string{
			"name": "sdjkhfajk",
		})
		md2 := metadata.Pairs("name", "jjjjj")*/
	/*md1 := metadata.New(map[string]string{
		"name": " sdjkhfajk",
	})*/
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
