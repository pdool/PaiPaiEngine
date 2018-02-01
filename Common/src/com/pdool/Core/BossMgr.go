package Core

import "com/pdool/Core/Data"

var guid int64 = 0

type BossMgr struct {
}

func CreateObjID() Data.GUID {
	guid++
	id := Data.GUID{AppID: 1, Value: guid}
	return id
}
