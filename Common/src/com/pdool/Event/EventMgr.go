package Core

import (
	"com/pdool/Core"
	"sync"
)

var eventMgrInstance *EventMgr
var lock  = &sync.Mutex {}

func GetInstance() *EventMgr {
	lock.Lock()
	defer lock.Unlock()
	if eventMgrInstance == nil {
		eventMgrInstance = &EventMgr {}
	}
	return eventMgrInstance
}

type EventMgr struct {
	msgQueue *Core.Queue
}
func(this EventMgr)New(){
	this.msgQueue = Core.NewQueue()

}

