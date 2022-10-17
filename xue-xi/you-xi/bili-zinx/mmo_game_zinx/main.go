package main

import (
	"bili-zinx/mmo_game_zinx/core"
	"bili-zinx/zinx/ziface"
	"bili-zinx/zinx/znet"
	"fmt"
)



func OnConnecionAdd(conn ziface.IConnection) {

	player := core.NewPlayer(conn)

	//同步当前的pid给客户端
	player.SyncPid()

	//同步当前的初始位置
	player.BroadCastStartPosition()

	fmt.Println("==OnConnecionAdd==")
}



func main() {
	//创建server句柄
	s := znet.NewService()



	s.Serve()
}
