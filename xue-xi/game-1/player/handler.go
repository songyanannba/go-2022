package player

import "game-1/function"

type Handler func(interface{})

func (p *Player) AddFriend(data interface{}) {

	fId := data.(uint64)
	if !function.CheckInNumberSlice(fId, p.FriendList) {
		p.FriendList = append(p.FriendList, fId)
	}
}
