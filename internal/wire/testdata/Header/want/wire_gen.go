// This is a sample header file.
//
// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/anhro/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectFoo() Foo {
	foo := provideFoo()
	return foo
}
