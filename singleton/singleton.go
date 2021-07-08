package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	client string
	// other things
}

var instance singleton = singleton{}

//func NewInstance() *singleton {
//	if instance!=nil {
//		return instance
//	}
//	instance = &singleton{client: "client"}
//	return instance
//}

var once sync.Once

func Init() {
	// this will only work once
	once.Do(func() {
		fmt.Println("initializing")
		instance.client = "client2"
	})
}

func main() {
	//s := NewInstance()
	Init()
	fmt.Println(instance.client)
	Init()
	fmt.Println(instance.client)
}
