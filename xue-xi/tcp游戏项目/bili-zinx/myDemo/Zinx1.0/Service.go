package main

import (
	"bili-zinx/zinx/ziface"
	"bili-zinx/zinx/znet"
	"fmt"
)

//自定义路由
type PingRouter struct {
	znet.BaseRouter
}




func (p *PingRouter) Handle(request ziface.IRequest) {

	fmt.Println("call ping router Handler...")

	//先读区客户端数据
	fmt.Println("recv from client :msgId = ", request.GetMsgID() , ",data ," , string(request.GetData()))

	err := request.GetConnection().SendMsg(200 ,[]byte("ping... ping...ping..."))

	if err != nil {
		fmt.Println(err)
	}

}


type HelloRouter struct {
	znet.BaseRouter
}


func DoConnBegin(conn ziface.IConnection) {
	fmt.Println("==>do conn begin is call...")
	if err := conn.SendMsg(202 , []byte("DoConnBegin...")) ; err != nil {
		fmt.Println(err)
	}
}

func DoConnLost(conn ziface.IConnection) {
	fmt.Println("==>do conn Lost is call...")
	//服务资源回收
	fmt.Println("service controller...ConnID " ,conn.GetConnID() , "is Lost")
}



func (p *HelloRouter) Handle(request ziface.IRequest) {

	fmt.Println("call hello router Handler...")

	//先读区客户端数据
	fmt.Println("recv from client :msgId = ", request.GetMsgID() , ",data ," , string(request.GetData()))

	err := request.GetConnection().SendMsg(201 ,[]byte("world... world...world..."))

	if err != nil {
		fmt.Println(err)
	}

}


func main() {

	//创建service句柄
	s := znet.NewService()

	s.SetOnConnStart(DoConnBegin)
	s.SetOnConnStop(DoConnLost)

	s.AddRouter(0 ,&PingRouter{})
	s.AddRouter(1 ,&HelloRouter{})

	//启动service
	s.Serve()

}
