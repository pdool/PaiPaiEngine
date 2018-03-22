package Core

type ModuleMgr struct {
	mgrs map[string]interface{}
}

var mgr *ModuleMgr

func GetModuleMgr() *ModuleMgr {
	if mgr == nil {
		mgr = &ModuleMgr{mgrs: make(map[string]interface{})}
	}
	return mgr
}

func (mgr ModuleMgr) RegistModule(moduleName string, modulePointer interface{}) {
	if _, ok := mgr.mgrs[moduleName]; ok {
		return
	}
	mgr.mgrs[moduleName] = modulePointer
}
func (mgr ModuleMgr) GetModule(moduleName string) interface{} {
	if module, ok := mgr.mgrs[moduleName]; ok {
		return module
	}
	return nil
}
