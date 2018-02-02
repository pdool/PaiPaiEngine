package Core

const(
	RECORD_ADD = iota
	RECORD_DEL
	RECORD_UPDATE
	RECORD_CREATE
)

type IRecordHandler interface {
	Handle(obj GUID, recordName string ,opType int)
}
