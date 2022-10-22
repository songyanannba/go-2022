package core

import (
	"fmt"
	"testing"
)

func TestNewAoiManager(t *testing.T) {
	AOIMgr := NewAoiManager(0, 250, 5, 0, 250, 5)

	fmt.Println(AOIMgr)
}

func TestAOIManager_GetSurroundGridsByGid(t *testing.T) {
	AOIMgr := NewAoiManager(0, 250, 5, 0, 250, 5)

	for gid, _ := range AOIMgr.grids {
		grids := AOIMgr.GetSurroundGridsByGid(gid)
		fmt.Println("gid :", gid)
		gIDs := make([]int, 0, len(grids))
		for _, grid := range grids {
			gIDs = append(gIDs, grid.GID)
		}
		fmt.Println("grid IDs are ", gIDs)
	}
}
