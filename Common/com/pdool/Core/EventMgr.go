package Core

import (
	"sync"
	"fmt"
	"com/pdool/DataStruct"
)

var eventMgrInstance *EventMgr
var eventLock = &sync.Mutex{}

func GetEventMgr() *EventMgr {
	eventLock.Lock()
	defer eventLock.Unlock()
	if eventMgrInstance == nil {
		eventMgrInstance = &EventMgr{msgQueue: &DataStruct.Queue{}}
	}
	return eventMgrInstance
}

type EventMgr struct {
	msgQueue *DataStruct.Queue
}

func (eventMgr *EventMgr) Start() {

	for eventMgr.msgQueue.Size() > 0 {
		ele := eventMgr.msgQueue.Dequeue()
		event := ele.(*Event)
		go dealEvent(event)
	}
}

func(eventMgr *EventMgr) Push(event *Event){
	eventMgr.msgQueue.Enqueue(event)
}

func dealEvent(event *Event) {
	name := event.GetComponentName()
	id := event.GetEventId()
	msg := event.GetMsg()
	fmt.Println(name)
	fmt.Println(id)
	fmt.Println(msg)
}
