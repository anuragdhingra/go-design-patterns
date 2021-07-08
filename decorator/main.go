package main

import "fmt"

type Beverage interface {
	GetCost() int
}

// Concrete components
type espresso struct{}
func (e espresso) GetCost() int {
	return 1
}

type cafelatte struct{}
func (c cafelatte) GetCost() int {
	return 2
}

// Decorators
type AddOnBeverageDecorator interface {
	GetCost() int
}

type creamMilk struct {
	beverage Beverage
}
func (c creamMilk) GetCost() int {
	return c.beverage.GetCost() + 10
}

type chocolate struct {
	beverage Beverage
}
func (c chocolate) GetCost() int {
	return c.beverage.GetCost() + 12
}

func main() {
	espresso := espresso{}
	espressoWithCream := creamMilk{beverage: espresso}
	fmt.Println("espressoWithCream price is", espressoWithCream.GetCost())
	espressoWithCreamAndChocolate := chocolate{beverage: espressoWithCream}
	fmt.Println("espressoWithCreamAndChocolate price is", espressoWithCreamAndChocolate.GetCost())
}
