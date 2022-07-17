package main

import "bili-zinx/zinx/znet"

func main( ) {

	//创建service句柄
	s := znet.NewService("[zinx 0.1]")

	//启动service
	s.Serve()

}
