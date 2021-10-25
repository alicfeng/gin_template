package event

import "github.com/alicfeng/gin_template/app/listener"

type DemoEvent struct {
	Event
}

// NewDemoEvent /**
func NewDemoEvent(payload interface{}) *DemoEvent {
	return (&DemoEvent{
		Event{
			Payload: payload,
		},
	}).bind()
}

// bind 绑定事件 /**
func (event *DemoEvent) bind() *DemoEvent {
	// a...
	event.AddListener(listener.NewDemo().Handle)

	return event
}
