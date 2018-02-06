package Core

import (
	"sync"
	"com/pdool/DataStruct"
)

type DataMgr struct {
	datas DataStruct.Dictionary
}

var dataMgrInstance *DataMgr
var dataLock = &sync.Mutex{}

func GetDataMgr() *DataMgr {
	dataLock.Lock()
	defer dataLock.Unlock()
	if dataMgrInstance == nil {
		dataMgrInstance = &DataMgr{}
	}
	return dataMgrInstance
}

func (dataMgr *DataMgr) HasData(key string) bool {
	return dataMgr.datas.Has(key)
}
func (dataMgr *DataMgr) GetData(key string) interface{} {
	return dataMgr.datas.Get(key)
}
func (dataMgr *DataMgr) SetData(key string, value interface{}) {
	dataMgr.datas.Set(key, value)
}
