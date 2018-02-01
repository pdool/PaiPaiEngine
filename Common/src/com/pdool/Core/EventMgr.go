package Core

import (
	"sync"
	"fmt"
	"com/pdool/Core/Event"
)

var eventMgrInstance *EventMgr
var eventLock = &sync.Mutex{}

func GetEventMgr() *EventMgr {
	eventLock.Lock()
	defer eventLock.Unlock()
	if eventMgrInstance == nil {
		eventMgrInstance = &EventMgr{msgQueue: Event.NewQueue()}
	}
	return eventMgrInstance
}

type EventMgr struct {
	msgQueue *Event.Queue
}

func (eventMgr *EventMgr) Start() {

	for eventMgr.msgQueue.Len() > 0 {
		ele := eventMgr.msgQueue.Pop()
		event := ele.Value.(*Event.Event)
		go dealEvent(event)
	}
}

func(eventMgr *EventMgr) Push(event *Event.Event){
	eventMgr.msgQueue.Push(event)
}

func dealEvent(event *Event.Event) {
	name := event.GetComponentName()
	id := event.GetEventId()
	msg := event.GetMsg()
	fmt.Println(name)
	fmt.Println(id)
	fmt.Println(msg)
}
