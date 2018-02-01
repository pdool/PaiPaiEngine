package Event

type Event struct {
	EventId       int
	ComponentName string
	MsgBody       interface{}
}

func (event Event) GetEventId() int {
	return event.EventId
}
func (event Event) GetComponentName() string {
	return event.ComponentName
}
func (event Event) GetMsg() interface{} {
	return event.MsgBody
}
