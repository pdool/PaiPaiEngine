package Core

var timeMgr *TimeMgr

func GetTimeMgr() *TimeMgr {
	if timeMgr == nil {
		timeMgr = &TimeMgr{}
	}
	return timeMgr
}

type TimeMgr struct {
}
