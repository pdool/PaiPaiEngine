package Core

type Event struct {
	eventId       int
	componentName string
	msg           interface{}
}

func (this Event) GetEventId() int {
	return this.eventId
}
func (this Event) GetComponentName() string {
	return this.componentName
}
func (this Event) GetMsg() interface{} {
	return this.msg
}
