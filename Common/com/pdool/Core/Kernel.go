package Core

var pKernel *Kernel
type Kernel struct {

	timeMgr TimeMgr			//	时间管理器
	sceneMgr SceneManager	//	场景管理器
	configMgr ConfigManager	//	配置管理器
	eventMgr EventMgr		//	事件管理器
}
func GetKernel() *Kernel {
	if pKernel == nil {
		pKernel = &Kernel{}
	}
	return pKernel
}
func (pKernel *Kernel)GetObject(guid GUID)*Object{

	return nil
}
