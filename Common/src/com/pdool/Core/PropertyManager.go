package Core

import "com/pdool/DataStruct"

type PropertyManager struct {
	guid  GUID
	props DataStruct.Dictionary
}

func NewPropertyManager(guid GUID) *PropertyManager {
	pMgr := new(PropertyManager)
	pMgr.guid = guid
	pMgr.props = DataStruct.Dictionary{}
	return pMgr
}
func (p *PropertyManager) AddProperty(guid GUID, propName string) {

}
func (p *PropertyManager) SetProperty(propName string, value interface{}) bool {
	prop := p.props.Get(propName).(*Property)
	if prop == nil {
		return false
	}
	prop.SetValue(value)
	return true
}
func (p *PropertyManager) HasProperty(propName string) bool {
	return p.props.Has(propName)
}
func (p *PropertyManager) GetPropertyInt(propName string) int {
	return p.props.Get(propName).(*Property).GetInt()
}
func (p *PropertyManager) GetPropertyFloat(propName string) float32 {
	return p.props.Get(propName).(*Property).GetFloat()
}
func (p *PropertyManager) GetPropertyString(propName string) string {
	return p.props.Get(propName).(*Property).GetString()
}
func (p *PropertyManager) GetPropertyObj(propName string) interface{} {
	return p.props.Get(propName).(*Property).GetObject()
}
