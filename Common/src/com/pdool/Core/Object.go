package Core


type Object struct {
	guid GUID
	className string
	propMgr PropertyManager
	recordMgr RecordManager
}
func NewObject(guid GUID){
	obj := new(Object)
	obj.guid = guid
}
func (p *Object)GetPropInt(prop string)int{
	return p.propMgr.GetPropertyInt(prop)
}


func (p *Object) SetPropertyValue(propName string, value interface{}) {
	p.propMgr.SetProperty(propName,value)
}