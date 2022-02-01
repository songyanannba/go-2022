package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mic-basic-lessons/week5/5-9/ptoto/pb"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())

	client := pb.NewStudyClient(conn)

	resp, _ := client.Study(context.Background(), &pb.BoolRequest{
		Name: "kkk",
	})
	fmt.Println(resp.Msg)
	fmt.Println("end....")
}
