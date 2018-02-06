package Core

import (
	"com/pdool/DataStruct"
	"com/pdool/Math"
)

type Scene struct {
	sceneId    int
	groups     DataStruct.Dictionary //	groupid : sceneGroup
	groupIndex int
	relivePos  DataStruct.LinkedList
}

func CreateScene(sceneId int) *Scene {
	scene := new(Scene)
	scene.groups = DataStruct.Dictionary{}
	scene.groupIndex = -1
	scene.relivePos = DataStruct.LinkedList{}
	return scene
}

//	创建一个分线
func (s *Scene) NewGroupId() int {
	s.groupIndex++
	return s.groupIndex
}

//	重生点
func (s *Scene) GetRebornPos() DataStruct.LinkedList {
	return s.relivePos
}
func (s *Scene) AddRebornPos(vector Math.Vector3) {
	s.relivePos.Append(vector)
}

//	分线中是否存在对象guid
func (s *Scene) ExistInGroup(groupId int, guid GUID) bool {
	sceneGroup := s.groups.Get(groupId)
	if sceneGroup != nil {
		group := sceneGroup.(*SceneGroup)
		if group.playerList.IndexOf(guid) != -1 || group.otherList.IndexOf(guid) != -1 {
			return true
		}
	}
	return false
}

//	从分线中移出对象
func (s *Scene) RemoveObjFromGroup(groupId int, guid GUID) bool {
	sceneGroup := s.groups.Get(groupId)
	if sceneGroup != nil {
		group := sceneGroup.(*SceneGroup)
		if group.playerList.IndexOf(guid) != -1 {

			group.playerList.Remove(guid)
			return true
		}
		if group.otherList.IndexOf(guid) != -1 {
			group.otherList.Remove(guid)
			return true
		}
	}
	return false
}

//	把对象加入到分线中
func (s *Scene) AddObjectToGroup(groupId int, guid GUID, bPlayer bool) bool {
	sceneGroup := s.groups.Get(groupId)
	if sceneGroup != nil {
		group := sceneGroup.(*SceneGroup)
		if bPlayer && group.playerList.IndexOf(guid) != -1 {
			group.playerList.Append(guid)
			return true
		} else if group.otherList.IndexOf(guid) != -1 {
			group.otherList.Append(guid)
			return true
		}
		return false
	}
	return false
}
