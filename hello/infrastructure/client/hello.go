package client

func NewHello() Hello {
	return &helloImpl{}
}

type Hello interface {
	Get() string
}

type helloImpl struct {}

func (c *helloImpl) Get() string {
	return "hello"
}
