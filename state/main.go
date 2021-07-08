package main

import (
	"fmt"
	"reflect"
)

//-------------------------------------------------------
type Gate struct {
	state GateState
}

func (g *Gate) PayOk() {
	//fmt.Println(reflect.TypeOf(g.state))
	g.state.PayOk()
}
func (g *Gate) PayFailed() { g.state.PayFailed() }
func (g *Gate) Pay() {
	//fmt.Println(reflect.TypeOf(g.state))
	g.state.Pay()
}
func (g *Gate) Enter() { g.state.Enter() }
func (g *Gate) ChangeState(state GateState) {
	fmt.Printf("Changing state from %v to %v \n", reflect.TypeOf(g.state), reflect.TypeOf(state))
	g.state = state
}
func NewGate() *Gate {
	g := &Gate{}
	g.state = NewClosedGateState(g)
	return g
}

//-------------------------------------------------------

type GateState interface {
	PayOk()
	PayFailed()
	Pay()
	Enter()
}

//-------------------------------------------------------
type OpenGateState struct {
	gate *Gate
}

func (g OpenGateState) PayOk() {
	// let user in
	g.gate.ChangeState(NewClosedGateState(g.gate))
}
func (g OpenGateState) PayFailed() {}
func (g OpenGateState) Pay() {
	fmt.Println("gate is open")
}
func (g OpenGateState) Enter() {}
func NewOpenGateState(g *Gate) GateState {
	return OpenGateState{
		gate: g,
	}
}

//-------------------------------------------------------
type ClosedGateState struct {
	gate *Gate
}

func (g ClosedGateState) PayOk() {
	// let user in
	g.gate.ChangeState(NewOpenGateState(g.gate))
}
func (g ClosedGateState) PayFailed() {}
func (g ClosedGateState) Pay() {
	fmt.Println("gate is closed")
}
func (g ClosedGateState) Enter() {}
func NewClosedGateState(g *Gate) GateState {
	return ClosedGateState{
		gate: g,
	}
}

//-------------------------------------------------------
func main() {
	gate := NewGate()
	gate.Pay()
	gate.PayOk()
	gate.Pay()
	gate.PayOk()
}
