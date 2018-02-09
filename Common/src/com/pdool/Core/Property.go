package Core

import (
	"container/list"
)

const (
	Prop_flag_self_public  = iota
	Prop_flag_other_public
)

type Property struct {
	guid         GUID        //	对象id
	name         string      //	属性名
	propType     int         //	属性类型
	selfFlag     int         //	自己是否可见属性
	otherFlag    int         //	其他人是否可见属性
	defaultValue interface{} //	默认值
	isSave       bool        //	是否保存

	value     interface{} //	值
	callbacks *list.List  //	回调
}

func NewProperty(guid GUID, name string, propType int, defaultValue interface{}, isSave bool) *Property {
	p := new(Property)
	p.guid = guid
	p.name = name
	p.propType = propType
	p.isSave = isSave
	return p
}

//	获取属性名
func (prop *Property) GetPropName() string {
	return prop.name
}

//	获取属性类型
func (prop *Property) GetPropType() int {
	return prop.propType
}

//	设置属性值
func (prop *Property) SetValue(value interface{}) {
	oldValue := prop.value
	newValue := value
	prop.value = value

	for cb := prop.callbacks.Front(); cb != nil; cb = cb.Next() {
		handler := cb.Value.(IPropHandler)
		handler.PropHandle(prop.guid, prop.name, oldValue, newValue)
	}

}

//	获取属性值
func (prop *Property) GetValue() (interface{}) {
	return prop.value
}

//	添加回调
func (prop *Property) AddCallBack(cb IPropHandler) {
	if prop.callbacks == nil {
		prop.callbacks = list.New()
	}
	prop.callbacks.PushBack(cb)
}
