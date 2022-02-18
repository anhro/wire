package main

import (
	"strings"
)

type (
	A interface {
		GetName() string
	}

	aImpl struct {
		b B
		c C
	}
)

func NewA(b B, c C) A {
	return &aImpl{
		b: b,
		c: c,
	}
}

func (a *aImpl) GetName() string {
	return strings.Join([]string{
		"A",
		a.b.GetName(),
		a.c.GetName(),
	}, ".")
}

//func BindA(a *aImpl) A {
//	return a
//}
