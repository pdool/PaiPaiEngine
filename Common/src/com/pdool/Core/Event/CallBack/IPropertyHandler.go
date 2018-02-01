package CallBack

import "com/pdool/Core/Data"

type IPropHandler interface {
	Handle(obj Data.GUID, propName string ,oldValue interface{},newValue interface{})
}
