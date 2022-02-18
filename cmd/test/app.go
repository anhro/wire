package main

import (
	"fmt"
)

type (
	App struct {
		a A
	}
)

func NewApp(a A) *App {
	app := &App{
		a: a,
	}

	return app
}

func (app *App) Start() error {
	if app.a.GetName() == "" {
		return fmt.Errorf("name of a is empty")
	}

	fmt.Println(app.a.GetName())

	return nil
}

//func BindApp(app *App) *App {
//	return app
//}
