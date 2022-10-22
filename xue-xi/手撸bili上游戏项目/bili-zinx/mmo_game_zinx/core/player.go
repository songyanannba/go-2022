package core

import (
	"bili-zinx/zinx/ziface"
	"fmt"
	"github.com/golang/protobuf/proto"
	"math/rand"
	"sync"
)

type Player struct {
	Pid  int32
	Conn ziface.IConnection
	X    float32
	Y    float32
	Z    float32
	V    float32
}

//player id 生成器
var PidGen int32 = 1
var IdLock sync.Mutex

func NewPlayer(conn ziface.IConnection) *Player {

	IdLock.Lock()
	id := PidGen
	id++
	IdLock.Unlock()

	return &Player{
		Pid:  1, //唯一主键
		Conn: conn,
		X:    float32(160 + rand.Intn(10)),
		Y:    0,
		Z:    float32(130 + rand.Intn(10)),
		V:    0,
	}
}

//发送给客户端消息
func (p *Player) SendMsg(magId uint32, data proto.Message) {



	msg, err := proto.Marshal(data)

	if err != nil {
		fmt.Println("player send mag maeshal err...")
		return
	}

	if p.Conn == nil {
		return
	}

	if err := p.Conn.SendMsg(magId, msg); err != nil {
		return
	}
	return
}

func (p *Player) SyncPid() {
	data := &pb.SyncPid{
		Pid: p.Pid,
	}
	p.SendMsg(1, data)
}

func (p *Player)BroadCastStartPosition(){
	protoc_msg := &pb.BroadCast {
		Pid : p.Pid,
		Tp : 2,
		Data: &pb.BroadCast_p{
			p :&pb.Position{
				X : p.X,
				Y : p.Y,
				Z : p.Z,
				Z :p.Z,
			},
		},
	}

	p.SendMsg(200, protoc_msg)
}