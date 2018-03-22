package Core

import "com/pdool/DataStruct"

type IPropHandler interface {
	PropHandle(obj GUID, propName string, oldValue interface{}, newValue interface{})
}

type IRecordHandler interface {
	RecordHandle(obj GUID, recordName string, opType int, row int, col int, oldValue DataStruct.LinkedList, newValue DataStruct.LinkedList)
}

type IBeforeLeaveSceneHandler interface {
	Handle(obj GUID, scenId int, groupId int, args interface{})
}
type IAfterLeaveSceneHandler interface {
	Handle(obj GUID, scenId int, groupId int, args interface{})
}

type IBeforeEnterSceneHandler interface {
	Handle(obj GUID, scenId int, groupId int, args interface{})
}
type IAfterEnterSceneHandler interface {
	Handle(obj GUID, scenId int, groupId int, args interface{})
}
