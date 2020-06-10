package foo

type Issuer interface {
	Get() Container
}

type Container interface {
	Get() string
}

type MyIssuer struct {
}

func (i MyIssuer) Get() Container {
	return &MyContainer{}
}

type MyContainer struct {
}

func (c MyContainer) Get() string {
	return "Hello, World!"
}
