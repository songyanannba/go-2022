package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"mic-basic-lessons/week5/5-12/proto/pb"
	"sync"
)

func main() {
	conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewFoodServiceClient(conn)

	//服务端流模式
	/*client := pb.NewFoodServiceClient(conn)
	resp, err := client.SayName(context.Background(), &pb.FoodStreamRequest{
		Name: "麻辣小龙虾1",
	})
	if err != nil {
		panic(err)
	}

	for {
		recv, err := resp.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(recv.Msg)
		fmt.Println("...end")
	}*/

	//客户端流模式
	/*ctx, cancelFunc := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancelFunc()
	pCliet, err := client.PostName(ctx)

	if err != nil {
		panic(err)
	}

	foods := []string{"饺子", "东坡肘子", "水煮牛肉", "既定"}

	for _, item := range foods {
		fmt.Println("上菜" + item)
		err := pCliet.Send(&pb.FoodStreamRequest{
			Name: item,
		})
		time.Sleep(time.Second * 1)
		if err != nil {
			fmt.Println(err)
			break
		}
	}*/

	foods := []string{"饺子", "东坡肘子", "水煮牛肉", "大盘鸡"}

	fullClient, err := client.FullStream(context.Background())
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(2)

	//收
	go func() {
		defer wg.Done()
		for {
			item, err2 := fullClient.Recv()
			if err2 != nil {
				fmt.Println(err2)
				break
			}
			fmt.Println(item.Msg)
		}
	}()

	go func(s []string) {
		defer wg.Done()
		for _, item := range s {
			err = fullClient.Send(&pb.FoodStreamRequest{Name: item})
			if err != nil {
				fmt.Println(err)
			}
		}
	}(foods)

	wg.Wait()

}
