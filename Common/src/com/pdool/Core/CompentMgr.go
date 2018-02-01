package Core

import (
	"sync"
	"log"
)

var instance *CompentMgr
var lock = &sync.Mutex{}

func GetCompentMgr() *CompentMgr {
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

/**
注册组件
 */
func (this CompentMgr) RegisterCompent(compentName string, compent *interface{}) {
	this.compents[compentName] = compent
}

/**
查找组件
 */
func (this CompentMgr) GetCompent(compentName string) *interface{} {
	v, ok := this.compents[compentName]
	if ok {
		return v
	} else {
		log.Fatal("没有找到组件！！")
		return nil
	}
}
