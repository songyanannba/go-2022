package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"mic-trainning-lessons/account_srv/biz"
	"mic-trainning-lessons/account_srv/proto/pb"
	"mic-trainning-lessons/internal"
	"mic-trainning-lessons/util"
	"net"
)

func init() {
	internal.InitDB()
}

func getAddr1() string {
	//获取默认地址的方法
	ip := flag.String("ip", "192.168.1.4", "输入IP")
	port := flag.Int("port", 9095, "输入端口")
	flag.Parse()
	return fmt.Sprintf("%s:%d", *ip, *port)
}

func main() {
	//addr := getAddr1()

	//改造 获取地址的方法
	port := util.GenRandomPort()
	addr := fmt.Sprintf("%s:%d", internal.AppConf.ConsulConfig.Host, port)

	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server, &biz.AccountServer{})
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		zap.S().Error("account_srv异常..." + err.Error())
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
		internal.AppConf.AccountSrvConfig.Host,
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
		Name:    internal.AppConf.AccountSrvConfig.SrvName,
		ID:      randUUID,
		Port:    port,
		Tags:    internal.AppConf.AccountSrvConfig.Tags,
		Address: internal.AppConf.AccountSrvConfig.Host,
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
