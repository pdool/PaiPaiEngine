package Core

import "com/pdool/DataStruct"

type PropertyManager struct {
	guid  GUID
	props DataStruct.Dictionary
}

//	生成一个属性管理器
func NewPropertyManager(guid GUID) *PropertyManager {
	pMgr := new(PropertyManager)
	pMgr.guid = guid
	pMgr.props = DataStruct.Dictionary{}
	return pMgr
}

//	添加一个属性
func (p *PropertyManager) AddProperty(guid GUID, propName string, propType int, defaultValue interface{}, isSave bool) bool {
	if p.HasProperty(propName) {
		return false
	}
	prop := NewProperty(guid, propName, propType, defaultValue, isSave)
	p.props.Set(propName, prop)
	return true
}

//	设置属性值
func (p *PropertyManager) SetProperty(propName string, value interface{}) bool {
	prop := p.props.Get(propName).(*Property)
	if prop == nil {
		return false
	}
	prop.SetValue(value)
	return true
}

//	属性是否存在
func (p *PropertyManager) HasProperty(propName string) bool {
	return p.props.Has(propName)
}

//	获取整型
func (p *PropertyManager) GetPropertyInt(propName string) int {
	return p.props.Get(propName).(*Property).GetValue().(int)
}
func (p *PropertyManager) GetPropertyFloat(propName string) float32 {
	return p.props.Get(propName).(*Property).GetValue().(float32)
}
func (p *PropertyManager) GetPropertyString(propName string) string {
	return p.props.Get(propName).(*Property).GetValue().(string)
}
func (p *PropertyManager) GetPropertyObj(propName string) interface{} {
	return p.props.Get(propName).(*Property).GetValue().(GUID)
}
