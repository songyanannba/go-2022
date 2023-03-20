package core

import (
	"fmt"
	"sync"
)

/**
	一个AOI格子的数据类型
 */
type Grid struct {
	//格子ID
	GID int
	//格子的左边坐标
	MinX int
	//格子的右边坐标
	MaxX int
	//格子的上边坐标
	MinY int
	//格子的下边坐标
	MaxY int
	//玩家ID
	PlayerIDs map[int]bool
	//锁
	PIdLock sync.RWMutex
}


func NewGrid(gId ,minX ,maxX , minY , maxY int) *Grid {
	return &Grid{
		GID:       gId,
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
		PlayerIDs: make(map[int]bool),
	}
}

func(g *Grid)Add(playerID int) {
	g.PIdLock.Lock()
	defer  g.PIdLock.Unlock()

	g.PlayerIDs[playerID] = true
}

func(g *Grid)Remove(playerID int) {
	g.PIdLock.Lock()
	defer  g.PIdLock.Unlock()

	delete(g.PlayerIDs , playerID)
}


func(g *Grid)GetPlayerIds() (playerIDs []int) {
	g.PIdLock.Lock()
	defer  g.PIdLock.Unlock()

	for k , _ := range g.PlayerIDs {
		playerIDs = append(playerIDs , k)
	}
	return
}


func(g *Grid) String() string {
	return fmt.Sprintf("Grid id :%d , minX is %d , manX is %d, minY is %d , manY is %d",
	g.GID , g.MaxX ,g.MaxX , g.MinY , g.MaxY)
}
