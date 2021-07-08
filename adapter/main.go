package main

import "fmt"

// ------------------Adaptee and adapter are coupled-----------------------------
type Adaptee1 struct{}

func (a Adaptee1) SpecificMethod() {
	fmt.Println("Specific Adaptee 1 method")
}

type Adapter1 struct{ Adaptee Adaptee1 } // has an adaptee
func (a Adapter1) Request() {
	a.Adaptee.SpecificMethod()
}

// -----------------------------------------------
type Adaptee2 struct{}

func (a Adaptee2) SpecificMethod() {
	fmt.Println("Specific Adaptee 2 method")
}

type Adapter2 struct {
	// has an adaptee
	Adaptee Adaptee2
}

func (a Adapter2) Request() {
	a.Adaptee.SpecificMethod()
}

// ---------------------Target client----------------------------
type TargetClient struct{}

func (t TargetClient) Anything(targetInterface TargetInterface) {
	targetInterface.Request()
}

// All adapters will implement Target interface
type TargetInterface interface {
	Request()
}

func main() {
	adapter1 := Adapter1{Adaptee: Adaptee1{}}
	adapter2 := Adapter2{Adaptee: Adaptee2{}}

	client := TargetClient{}
	client.Anything(adapter1)
	client.Anything(adapter2)
}
