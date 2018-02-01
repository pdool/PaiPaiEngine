package Data

type Object struct {
	guid GUID
	propMgr PropertyManager
	recordMgr RecordManager
}
func NewObject(guid GUID){
	obj := new(Object)
	obj.guid = guid
}
