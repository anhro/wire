//go:build wireinject
// +build wireinject

package main

import (
	"github.com/anhro/wire"
)

func BuildCInCompileTimeWithWire() (C, error) {
	wire.Build(
		NewC,
		NewD,
	)

	return nil, nil
}
