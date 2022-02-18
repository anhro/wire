package main

import (
	"fmt"
)

func main() {
	runApp()
	runA()
	runB()
	runC()

	fmt.Println("Finish")
}

func runApp() {
	//app := BuildAppInCompileTimeWithWire()
	//
	//err := app.Start()
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println("Finish App")
}

func runA() {
	//a, err := BuildAInCompileTimeWithWire()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("a.name %s\n", a.GetName())
	fmt.Println("Finish A")
}

func runB() {
	//b := BuildBInCompileTimeWithWire()
	//
	//fmt.Printf("b.name %s\n", b.GetName())
	fmt.Println("Finish B")
}

func runC() {
	//c := BuildCInCompileTimeWithWire()
	//
	//fmt.Printf("c.name %s\n", c.GetName())
	fmt.Println("Finish C")
}
