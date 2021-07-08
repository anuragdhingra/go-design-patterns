package main

import (
	"fmt"
	"math/rand"
)

type Animal interface {
	GetName() string
}

type duck struct {}
func (d duck) GetName() string {
	return "duck"
}

type dog struct {}
func (d dog) GetName() string {
	return "dog"
}

type bird struct {}
func (b bird) GetName() string {
	return "bird"
}

type AnimalFactory interface {
	Create() Animal
}

type randomAnimalFactory struct {}
func (r randomAnimalFactory) Create() Animal {
	i := rand.Int31n(3)
	if i == 0 {
		return duck{}
	} else if i == 1 {
		return dog{}
	} else {
		return bird{}
	}
}

func main() {
	randomAnimalFactory := randomAnimalFactory{}
	for i := 0; i<10 ;i++ {
		fmt.Println(randomAnimalFactory.Create().GetName())
	}
}