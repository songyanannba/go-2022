package mian

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"mic-trainning-lessons-part4/biz"
	"mic-trainning-lessons-part4/internal"
	"mic-trainning-lessons-part4/proto/pb"
	"net"
)

func init() {
	internal.InitDB()
}


func main() {

	//改造 获取地址的方法
	//port := util.GenRandomPort()
	port := internal.AppConf.CartOrderSrvConfig.Port
	addr := fmt.Sprintf("%s:%d", internal.AppConf.CartOrderSrvConfig.Host, port)

	server := grpc.NewServer()


	pb.RegisterShopCarServiceServer(server, &biz.ShopCartService{})

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		zap.S().Error("stock_srv异常..." + err.Error())
		panic(err)
	}
	fmt.Println("suc...")

	//grpc 健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	defaultConfig := api.DefaultConfig()
	defaultConfig.Address = fmt.Sprintf("%s:%d",
		internal.AppConf.ConsulConfig.Host,
		internal.AppConf.ConsulConfig.Port)

	client, err := api.NewClient(defaultConfig)
	if err != nil {
		panic(err)
	}

	checkAddr := fmt.Sprintf("%s:%d",
		internal.AppConf.CartOrderSrvConfig.Host,
		port,
	)
	check := &api.AgentServiceCheck{
		GRPC:                           checkAddr,
		Timeout:                        "3s",
		Interval:                       "1s",
		DeregisterCriticalServiceAfter: "5s",
	}
	randUUID := uuid.New().String()
	reg := api.AgentServiceRegistration{
		Name:    internal.AppConf.CartOrderSrvConfig.SrvName,
		ID:      randUUID,
		Port:    port,
		Tags:    internal.AppConf.CartOrderSrvConfig.Tags,
		Address: internal.AppConf.CartOrderSrvConfig.Host,
		Check:   check,
	}
	err = client.Agent().ServiceRegister(&reg)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("uuod:%s==>port:%d", randUUID, port))
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
