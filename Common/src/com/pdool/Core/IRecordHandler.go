package Core

import "com/pdool/DataStruct"

const(
	RECORD_ADD = iota
	RECORD_DEL
	RECORD_UPDATE
	RECORD_CREATE
)

type IRecordHandler interface {
	Handle(obj GUID, recordName string ,opType int,row int ,col int,oldValue DataStruct.LinkedList,newValue DataStruct.LinkedList)
}
