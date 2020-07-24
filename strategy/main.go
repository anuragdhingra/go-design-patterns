package main

import "fmt"

// QuackStrategy interface
type QuackStrategy interface {
    quack()
}

// SimpleQuackStrategy impl for QuackStrategy
type SimpleQuackStrategy struct {}
func (SimpleQuackStrategy) quack() {
    fmt.Print("Simple quack strategy, ")
}

// NoQuackStrategy impl for QuackStrategy
type NoQuackStrategy struct {}
func (NoQuackStrategy) quack() {
    fmt.Print("No quack strategy, ")
}

// DisplayStrategy interface
type DisplayStrategy interface {
    display()
}

// GraphicDisplayStrategy impl for DisplayStrategy
type GraphicDisplayStrategy struct {}
func (GraphicDisplayStrategy) display() {
    fmt.Print("Graphic display strategy, ")
}

// TextDisplayStrategy impl for DisplayStrategy
type TextDisplayStrategy struct {}
func (TextDisplayStrategy) display() {
    fmt.Print("Text display strategy, ")
}

// FlyStrategy interface
type FlyStrategy interface {
    fly()
}

// SimpleFlyStrategy impl for FlyStrategy
type SimpleFlyStrategy struct {}
func (sfs SimpleFlyStrategy) fly() {
    fmt.Println("Simple fly strategy")
}

// NoFlyStrategy impl for FlyStrategy
type NoFlyStrategy struct {}
func (nfs NoFlyStrategy) fly() {
    fmt.Println("No fly strategy")
}

// Duck struct
type Duck struct {
    QuackStrategy
    DisplayStrategy
    FlyStrategy
}

func main() {
    rubberDuck := Duck{NoQuackStrategy{}, GraphicDisplayStrategy{}, NoFlyStrategy{}}
    fmt.Print("rubberDuck: ")
    rubberDuck.QuackStrategy.quack()
    rubberDuck.DisplayStrategy.display()
    rubberDuck.FlyStrategy.fly()

    wildDuck :=  Duck{SimpleQuackStrategy{}, TextDisplayStrategy{}, SimpleFlyStrategy{}}
    fmt.Print("wildDuck: ")
    wildDuck.QuackStrategy.quack()
    wildDuck.DisplayStrategy.display()
    wildDuck.FlyStrategy.fly()

    cityDuck := Duck{SimpleQuackStrategy{}, GraphicDisplayStrategy{}, SimpleFlyStrategy{}}
    fmt.Print("cityDuck: ")
    cityDuck.QuackStrategy.quack()
    cityDuck.DisplayStrategy.display()
    cityDuck.FlyStrategy.fly()
}
