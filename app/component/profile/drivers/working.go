package profile

type Working struct {
	DemoMode int `json:"demo_mode"`
}

const (
	/*工作演示模式*/
	workingDemoModeEnable  = 1 // 启用
	workingDemoModeDisable = 2 // 停用
)

// WorkingDemoModeEnable 已经启动演示模式 /**
func (working Working) WorkingDemoModeEnable() bool {
	if workingDemoModeEnable == working.DemoMode {
		return true
	}
	return false
}
