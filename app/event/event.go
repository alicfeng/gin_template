package event

// Listeners 监听容器
type Listeners []ListenFunc

// ListenFunc 定义事件监听处理函数类型
type ListenFunc func(interface{})

// IEvent 事件接口
type IEvent interface {
	bind() *Event
}

// Event 事件结构体
type Event struct {
	Listeners Listeners
	Payload   interface{}
}

// AddListener 添加监听器 /**
func (event *Event) AddListener(listenFunc ListenFunc) *Event {
	event.Listeners = append(event.Listeners, listenFunc)

	return event
}

// Dispatch 事件调度分发 协程执行 /**
func (event *Event) Dispatch() {
	for _, item := range event.Listeners {
		go item(event.Payload)
	}
}
