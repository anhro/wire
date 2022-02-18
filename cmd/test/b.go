package main

import (
	"strings"
)

type (
	B interface {
		GetName() string
	}

	bImpl struct {
		c C
	}
)

func NewB(c C) B {
	return &bImpl{
		c: c,
	}
}

func (b *bImpl) GetName() string {
	return strings.Join([]string{
		"B",
		b.c.GetName(),
	}, ".")
}

//func BindB(b *bImpl) B {
//	return b
//}
