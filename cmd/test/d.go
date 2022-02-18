package main

type (
	D interface {
		GetName() string
	}

	dImpl struct{}
)

func NewD() D {
	return &dImpl{}
}

func (d *dImpl) GetName() string {
	return "D"
}

//func BindC(c *cImpl) C {
//	return c
//}
