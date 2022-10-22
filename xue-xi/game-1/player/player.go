package player

import (
	"game-1/chat"
	"game-1/function"
)

type Player struct {
	Uid        uint64
	FriendList []uint64 //朋友
	chChat     chan chat.Private
	Handlers   map[string]Handler
}

func NewPlayer() *Player {
	return &Player{
		Uid:        0,
		FriendList: make([]uint64, 100),
		Handlers:   make(map[string]Handler),
	}
}

func (p *Player) DelFriend(fId uint64) {
	p.FriendList = function.DelEleInSlice(fId, p.FriendList)
}

func (p *Player) Run() {

	for {
		select {
		case chatMsg := <-p.chChat:
			p.ResolveChatMsg(chatMsg)
		}
	}

}

func (p *Player) ResolveChatMsg(chatMsg chat.Private) {

}
