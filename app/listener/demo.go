package listener

type DemoListener struct {
}

func NewDemo() (listener *DemoListener) {
	return &DemoListener{}
}

func (listener *DemoListener) Handle(payload interface{}) {

}
