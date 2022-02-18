package main

import (
	"strings"
)

type (
	C interface {
		GetName() string
	}

	cImpl struct {
		d D
	}
)

func NewC(d D) (C, error) {
	result := &cImpl{
		d: d,
	}

	return result, nil
}

//func NewC() C {
//	return &cImpl{}
//}

func (c *cImpl) GetName() string {
	return strings.Join([]string{
		"C",
		c.d.GetName(),
	}, ".")
}

//func (c *cImpl) GetName() string {
//	return "C"
//}

//func BindC(c *cImpl) C {
//	return c
//}
