package manager

import "game-1/player"

//用户管理
type PlayerMgr struct {
	player map[uint64]player.Player
	addPCh chan player.Player
}

func (pm *PlayerMgr) Add(p player.Player) {
	pm.player[p.Uid] = p
}

func (pm *PlayerMgr) Run() {
	for {
		select {
		case p := <-pm.addPCh:
			pm.Add(p)
		}
	}
}
