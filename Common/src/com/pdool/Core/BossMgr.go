package Core


var guid int64 = 0

type BossMgr struct {
}

func CreateObjID() GUID {
	guid++
	id := GUID{AppID: 1, Value: guid}
	return id
}
