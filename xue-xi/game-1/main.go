package game_1

import "game-1/world"

func main() {
	world.MM = world.NewMgrMgr()
	world.MM.Pm.Run()

}
