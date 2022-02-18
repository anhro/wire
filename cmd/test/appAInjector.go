//go:build wireinject
// +build wireinject

package main

import (
	"github.com/anhro/wire"
)

func BuildAInCompileTimeWithWire() (A, error) {
	wire.Build(
		NewA,
		NewB,
		NewC,
	)

	return nil, nil
}
