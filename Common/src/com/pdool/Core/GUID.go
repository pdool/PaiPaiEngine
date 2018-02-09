package Core
var guidIndex int64 = 0
type GUID struct {

	AppID int
	Value int64
}

func NewGuid() *GUID {
	guid := new(GUID)
	guid.AppID = 1
	guidIndex ++
	guid.Value = guidIndex
	return guid
}

func (p *GUID)Equal(otherGuid *GUID)bool{
	if otherGuid!= nil && otherGuid.Value ==  p.Value && otherGuid.AppID == p.AppID{
		return true
	}
	return false
}
