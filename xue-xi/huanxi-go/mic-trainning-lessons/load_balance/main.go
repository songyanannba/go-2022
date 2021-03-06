package main

import (
	"context"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mic-trainning-lessons/account_srv/proto/pb"
	"mic-trainning-lessons/internal"
)


//负载均衡测试
func main() {
	addr := fmt.Sprintf("%s:%d", internal.AppConf.ConsulConfig.Host, internal.AppConf.ConsulConfig.Port)
	dialAddr := fmt.Sprintf("consul://%s/%s?wait=14", addr, internal.AppConf.AccountSrvConfig.SrvName)
	conn, err := grpc.Dial(dialAddr, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
	if err != nil {
		zap.S().Fatal(err)
	}
	defer conn.Close()

	client := pb.NewAccountServiceClient(conn)
	for i := 0; i < 10; i++ {
		res, err := client.GetAccountList(context.Background(), &pb.PagingRequest{
			PageNo:   1,
			PageSize: 3,
		})
		if err != nil {
			zap.S().Fatal(err)
		}

		for idx, item := range res.AccountList {
			fmt.Println(fmt.Sprintf("%d==>%s", idx, item))
		}
	}
}
