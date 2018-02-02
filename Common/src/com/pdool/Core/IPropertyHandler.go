package Core

type IPropHandler interface {
	Handle(obj GUID, propName string ,oldValue interface{},newValue interface{})
}
