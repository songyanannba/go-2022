package main

import (
	"bili-zinx/zinx/zifare"
	"bili-zinx/zinx/znet"
	"fmt"
	"math/rand"
)

//自定义路由
type PingRouter struct {
	znet.BaseRouter
}




func (p *PingRouter) PreHandle(request zifare.IRequest) {
	fmt.Println("call router PreHandler...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("PreHandler ping...\n"))
	if err != nil {
		fmt.Println("call back before ping error")
	}
}

func (p *PingRouter) Handle(request zifare.IRequest) {
	fmt.Println("call router Handler...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("Handler ping...\n"))
	if err != nil {
		fmt.Println("call back Handler ping error")
	}
}

func (p *PingRouter) PostHandle(request zifare.IRequest) {
	fmt.Println("call router PostHandler...")
	fmt.Println(rand.Int())
	fmt.Println("")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("PostHandler ping...\n"))
	if err != nil {
		fmt.Println("call back PostHandler ping error")
	}
}

func main() {

	//创建service句柄
	s := znet.NewService("[zinx 0.3]")

	s.AddRouter(&PingRouter{})

	//启动service
	s.Serve()

}
