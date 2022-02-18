//go:build wireinject
// +build wireinject

package main

import (
	"github.com/anhro/wire"
)

func BuildAppInCompileTimeWithWire() (*App, error) {
	wire.Build(
		NewA,
		NewB,
		NewC,
		NewApp,
	)

	return nil, nil
}
