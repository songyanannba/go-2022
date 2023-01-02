package client

import (
	"fmt"
	"game-1/network"
	"game-1/network/protocol/gen/player"
)

func (c *Client)OnLoginResp(pack *network.ClientPacket){
	rsp := &player.CSLogin{

	}


	fmt.Println("登陆成功。。。" ,rsp)
}
