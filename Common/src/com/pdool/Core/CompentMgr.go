package Core

import (
	"sync"
	"log"
)

var instance *CompentMgr
var lock = &sync.Mutex{}

func GetInstance() *CompentMgr {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = &CompentMgr{compents: make(map[string]*interface{})}
	}
	return instance
}

type CompentMgr struct {
	compents map[string]*interface{}
}

func (this CompentMgr) RegisterCompent(compentName string, compent *interface{}) {
	this.compents[compentName] = compent
}

func (this CompentMgr) GetCompent(compentName string) *interface{} {
	v, ok := this.compents[compentName]
	if ok {
		return v
	} else {
		log.Fatal("没有找到组件！！")
		return nil
	}
}
