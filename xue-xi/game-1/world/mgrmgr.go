package world

import "game-1/manager"

type MgrMgr struct {
	Pm manager.PlayerMgr
}

var MM *MgrMgr

func NewMgrMgr() *MgrMgr {
	m := &MgrMgr{Pm: manager.PlayerMgr{}}
	return m
}
