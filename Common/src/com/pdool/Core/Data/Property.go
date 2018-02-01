package Data

import (
	"container/list"
	"com/pdool/Core/Event/CallBack"
)

type Property struct {
	guid      GUID
	name      string
	value     Data
	oldValue  Data
	newValue  Data
	callbacks *list.List
}

/**

 */
func NewProperty2(guid GUID, name string) *Property {
	p := new(Property)
	p.guid = guid
	p.name = name
	return p
}
func NewProperty3(guid GUID, name string, value interface{}) *Property {
	p := new(Property)
	p.guid = guid
	p.name = name
	p.value = Data{value: value}
	return p
}
func (prop *Property) GetName() string {
	return prop.name
}

func (prop *Property) SetValue(value interface{}) {
	prop.oldValue = prop.value
	prop.newValue = Data{value: value}
	prop.value.SetData(value)

	for cb := prop.callbacks.Front(); cb != nil; cb = cb.Next() {
		handler := cb.Value.(CallBack.IPropHandler)
		handler.Handle(prop.guid,prop.name,prop.oldValue,prop.newValue)
	}
}

func (prop *Property) GetInt() int {

	return prop.value.GetInt()
}
func (prop *Property) GetFloat() float32 {
	return prop.value.GetFloat()
}
func (prop *Property) GetString() string {

	return prop.value.GetString()
}
func (prop *Property) GetObject() interface{} {
	return prop.value.GetObject()
}

func (prop *Property) AddCallBack(cb CallBack.IPropHandler) {
	if prop.callbacks == nil {
		prop.callbacks = list.New()
	}
	prop.callbacks.PushBack(cb)
}
