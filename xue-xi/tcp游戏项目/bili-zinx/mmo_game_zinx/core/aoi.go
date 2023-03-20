package core

import "fmt"

type AOIManager struct {
	MinX int

	MaxX int

	CntsX int

	MinY int

	MaxY int

	CntsY int

	grids map[int]*Grid
}

func NewAoiManager(minX, MaxX, cntsX, minY, maxY, cntsY int) *AOIManager {
	aoiMgr := &AOIManager{
		MinX:  minX,
		MaxX:  MaxX,
		CntsX: cntsX,
		MinY:  minY,
		MaxY:  maxY,
		CntsY: cntsY,
		grids: make(map[int]*Grid),
	}

	//格子的编号 的 初始化
	for y := 0; y < cntsY; y++ {
		for x := 0; x < cntsX; x++ {
			gid := y*cntsX + x
			aoiMgr.grids[gid] = NewGrid(
				gid,
				aoiMgr.MinX+x*aoiMgr.gridWidth(),
				aoiMgr.MinX+(x+1)*aoiMgr.gridWidth(),
				aoiMgr.MinY+y*aoiMgr.gridLength(),
				aoiMgr.MinY+(y+1)*aoiMgr.gridLength())
		}
	}

	return aoiMgr

}

func (m *AOIManager) gridWidth() int {
	return (m.MaxX - m.MinX) / m.CntsX
}

func (m *AOIManager) gridLength() int {
	return (m.MaxY - m.MinY) / m.CntsY
}

func (m *AOIManager) String() string {
	s := fmt.Sprintf("AOIManager :\n minX:%d, MaxX:%d, cntsX:%d, minY:%d, maxY:%d, cntsY:%d \n Grids in AOIManager :\n",
		m.MaxX, m.MaxX, m.CntsX, m.MinY, m.MaxY, m.CntsY)
	for _, grid := range m.grids {
		s += fmt.Sprintln(grid)
	}
	return s
}

//根据格子GID获取周围格子的ID集合

func (m *AOIManager) GetSurroundGridsByGid(gID int) (grids []*Grid) {

	//判断gID 是否在 AOIManger中
	if _, ok := m.grids[gID]; !ok {
		return nil
	}

	//初始化 grids 返回值切片
	grids = append(grids, m.grids[gID])

	//X 坐标 idX = id % nx
	idx := gID % m.CntsX

	//左边是否有格子
	if idx > 0 {
		grids = append(grids, m.grids[gID-1])
	}

	if idx < m.CntsX-1 {
		grids = append(grids, m.grids[gID+1])
	}

	//判断上下是否有格子
	gidsX := make([]int, len(grids))

	for _, v := range grids {
		gidsX = append(gidsX, v.GID)
	}

	for _, v := range gidsX {
		idy := v / m.CntsY
		if idy > 0 {
			grids = append(grids, m.grids[v-m.CntsX])
		}
		if idy < m.CntsY-1 {
			grids = append(grids, m.grids[v+m.CntsX])
		}

	}

	return
}

func (m *AOIManager) GetGidByPos(x, y float32) int {

	idx := (int(x) - m.MinX) / m.gridWidth()

	idy := (int(y) - m.MinY) / m.gridLength()

	return idy*m.CntsX + idx

}

func (m *AOIManager) GetPidByPos(x, y float32) (playerIDs []int) {

	gID := m.GetGidByPos(x, y)

	grids := m.GetSurroundGridsByGid(gID)

	for _, v := range grids {
		playerIDs = append(playerIDs, v.GetPlayerIds()...)
		fmt.Println(">>>")
	}
	return playerIDs
}

func (m *AOIManager) AddPidToGrid(pID, gID int) {
	m.grids[gID].Add(pID)
}

func (m *AOIManager) RemovePidFromGrid(pID, gID int) {
	m.grids[gID].Remove(pID)
}

func (m *AOIManager) GetPidFromGrid(gID int) (playerIDs []int) {
	playerIDs = m.grids[gID].GetPlayerIds()
	return nil
}

func (m *AOIManager) AddPidToGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	grid := m.grids[gID]
	grid.Add(pID)
}

func (m *AOIManager) RemoveFromGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	grid := m.grids[gID]
	grid.Remove(gID)
}
