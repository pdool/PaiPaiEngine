package Core

import (
	"com/pdool/Math"
	"com/pdool/Constant"
	"com/pdool/DataStruct"
	"strconv"
	"container/list"
	"fmt"
)

var smMgr *SceneManager

type SceneManager struct {
	scenes DataStruct.Dictionary

	//	todo 考虑去重问题
	beforeLeaveSceneCBs *list.List
	afterLeaveSceneCBs  *list.List
	beforeEnterSceneCBs *list.List
	afterEnterSceneCBs  *list.List
}

func GetSceneMgr() *SceneManager {
	if smMgr == nil {
		smMgr = new(SceneManager)
		smMgr.scenes = DataStruct.Dictionary{}
		smMgr.beforeLeaveSceneCBs = list.New()
		smMgr.afterLeaveSceneCBs = list.New()
		smMgr.beforeEnterSceneCBs = list.New()
		smMgr.afterEnterSceneCBs = list.New()

		smMgr.Init()
	}
	return smMgr
}
func (sm *SceneManager) Init() {
	fmt.Println("init")
	sm.scenes.Set(1,CreateScene(1))
}
func (sm *SceneManager) GetScene(sceneId int) *Scene{
	return sm.scenes.Get(sceneId).(*Scene)
}
func (sm *SceneManager) SwitchScene(self GUID, targetSceneId int, groupId int, pos Math.Vector3, orient float32, args interface{}) bool {
	obj := GetKernel().GetObject(self)
	if obj != nil {
		oldSceneId := obj.GetPropInt(Constant.Player_prop_sceneId)
		oldGroupId := obj.GetPropInt(Constant.Player_prop_groupId)

		if !sm.scenes.Has(oldSceneId) {
			Error(" no this container  " + strconv.Itoa(oldSceneId))
			return false
		}
		oldScene := sm.scenes.Get(oldSceneId).(*Scene)
		if !sm.scenes.Has(targetSceneId) {
			Error(" no this container " + strconv.Itoa(targetSceneId))
			return false
		}
		targetScene := sm.scenes.Get(targetSceneId).(*Scene)
		if !targetScene.groups.Has(groupId) {
			Error(" no this group " + strconv.Itoa(targetSceneId))
			return false
		}

		sm.beforeLeaveSceneGroup(self, oldSceneId, oldSceneId, args)
		oldScene.RemoveObjFromGroup(oldGroupId, self)
		obj.SetPropertyValue(Constant.Player_prop_groupId, 0)
		sm.afterLeaveSceneGroup(self, oldSceneId, oldSceneId, args)
		obj.SetPropertyValue(Constant.Player_prop_sceneId, targetSceneId)

		obj.SetPropertyValue(Constant.Player_prop_bornPos, targetScene.relivePos.Get(0))

		////////
		sm.beforeEnterSceneGroup(self, targetSceneId, groupId, args)

		targetScene.AddObjectToGroup(groupId, self, true)
		obj.SetPropertyValue(Constant.Player_prop_groupId, groupId)

		/////////
		sm.afterEnterSceneGroup(self, targetSceneId, groupId, args)

		return true
	}
	return false
}

//	离开场景之前
func (sm *SceneManager) beforeLeaveSceneGroup(self GUID, oldSceneId int, oldGroupId int, args interface{}) {
	for cb := sm.beforeLeaveSceneCBs.Front(); cb != nil; cb = cb.Next() {
		handler := cb.Value.(IBeforeLeaveSceneHandler)
		sm.AddBeforeLeaveCB(handler)
		handler.Handle(self, oldSceneId, oldGroupId, args)
	}
}

//	离开场景之后
func (sm *SceneManager) afterLeaveSceneGroup(self GUID, oldSceneId int, oldGroupId int, args interface{}) {
	for cb := sm.afterLeaveSceneCBs.Front(); cb != nil; cb = cb.Next() {
		handler := cb.Value.(IAfterLeaveSceneHandler)
		handler.Handle(self, oldSceneId, oldGroupId, args)
	}
}

//	进入场景之前
func (sm *SceneManager) beforeEnterSceneGroup(self GUID, sceneId int, groupId int, args interface{}) {
	for cb := sm.beforeEnterSceneCBs.Front(); cb != nil; cb = cb.Next() {
		handler := cb.Value.(IBeforeEnterSceneHandler)
		handler.Handle(self, sceneId, groupId, args)
	}
}

//	进入场景之后
func (sm *SceneManager) afterEnterSceneGroup(self GUID, sceneId int, groupId int, args interface{}) {
	for cb := sm.afterLeaveSceneCBs.Front(); cb != nil; cb = cb.Next() {
		handler := cb.Value.(IAfterEnterSceneHandler)
		handler.Handle(self, sceneId, groupId, args)
	}
}

//===================================callback===========================================================================

//	添加离开场景之前的回调
func (sm *SceneManager) AddBeforeLeaveCB(handler IBeforeLeaveSceneHandler) {
	sm.beforeLeaveSceneCBs.PushBack(handler)
}

//	添加离开场景之后的回调
func (sm *SceneManager) AddAfterLeaveCB(handler IBeforeLeaveSceneHandler) {
	sm.afterLeaveSceneCBs.PushBack(handler)
}

//	添加离开场景之前的回调
func (sm *SceneManager) AddBeforeEnterCB(handler IBeforeLeaveSceneHandler) {
	sm.beforeEnterSceneCBs.PushBack(handler)
}

//	添加离开场景之前的回调
func (sm *SceneManager) AddAfterEnterCB(handler IBeforeLeaveSceneHandler) {
	sm.afterEnterSceneCBs.PushBack(handler)
}

//===================================broadcast==========================================================================
// 	通知自己的数据更改()
//int OnObjectListEnter(const NFDataList& self, const NFDataList& argVar);
//int OnObjectListEnterFinished(const NFDataList& self, const NFDataList& argVar);
//
//int OnObjectListLeave(const NFDataList& self, const NFDataList& argVar);
//
////broad the data of self to argvar
//int OnPropertyEnter(const NFDataList& argVar, const NFGUID& self);
//int OnRecordEnter(const NFDataList& argVar, const NFGUID& self);
//
//int OnPropertyEvent(const NFGUID& self, const std::string& strProperty, const NFData& oldVar, const NFData& newVar, const NFDataList& argVar);
//int OnRecordEvent(const NFGUID& self, const std::string& strRecord, const RECORD_EVENT_DATA& xEventData, const NFData& oldVar, const NFData& newVar, const NFDataList& argVar);
//
//
//

//	属性回调
func (sm *SceneManager) PropHandle(obj GUID, propName string, oldValue interface{}, newValue interface{}) {

}
