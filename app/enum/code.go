package enum

// CodeEnum 业务编码统一引用结构体/**
type CodeEnum struct {
	Code    int
	Message string
}

const (
	SuccessCode = 1000
	FailureCode = 9999
)

/**
业务编码码表统一注册
*/
var (
	SUCCESS = CodeEnum{SuccessCode, "操作成功"}
	FAILURE = CodeEnum{FailureCode, "操作失败"}

	/*用户体系*/
	PasswordValidate = CodeEnum{FailureCode, "密码不正确"}

	/*设备体系*/
	DeviceUnbindBeforeDelete = CodeEnum{FailureCode, "删除设备前请解绑"}
	DeviceCaptureBeforeBind  = CodeEnum{FailureCode, "请绑定实例抓图服务"}

	/*服务模块*/
	// 抓图服务
	ServerSnapIsRunning            = CodeEnum{FailureCode, "抓图服务运行中,请停止再操作"}
	ServerSnapCommunicationFailure = CodeEnum{FailureCode, "抓图服务请求通信失败"}
	// 存储服务
	ServerMinioCommunicationFailure = CodeEnum{FailureCode, "储存服务请求失败"}
	// PLC
	ServerPLCCommunicationFailure = CodeEnum{FailureCode, "PLC服务请求失败"}

	// 报警记录
	AlarmRecordExistExport = CodeEnum{FailureCode, "存在任务正在导出中,请稍后再试"}
	AlarmRecordEmptyExport = CodeEnum{FailureCode, "查询记录为空,请重试"}
)
