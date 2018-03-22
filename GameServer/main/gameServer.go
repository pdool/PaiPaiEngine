package main

import (
	"fmt"
	"com/pdool/Core"
)


func main() {
	mgr := Core.GetModuleMgr()
	mgr.RegistModule("TimeMgr",Core.GetTimeMgr())
	mgr.RegistModule("SceneMgr",Core.GetSceneMgr())
	//mgr.RegistModule("FriendModule",new(TestHandler))
	manager := mgr.GetModule("SceneMgr").(*Core.SceneManager)

	scene := manager.GetScene(1)
	scene.NewGroupId()
	fmt.Println()
}
