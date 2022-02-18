//go:build wireinject
// +build wireinject

package main

import (
	"github.com/anhro/wire"
)

func BuildBInCompileTimeWithWire() (B, error) {
	wire.Build(
		NewB,
		NewC,
	)

	return nil, nil
}
