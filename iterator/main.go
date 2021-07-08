package main

import "fmt"

type Inventory interface {
	GetIterator() InventoryIterator
}

type HandHeldInventory struct {
	LeftHandItem  int
	RightHandItem int
}

func NewHandHeldInventory(LeftHandItem, RightHandItem int) Inventory {
	return HandHeldInventory{
		LeftHandItem:  LeftHandItem,
		RightHandItem: RightHandItem,
	}
}
func (h HandHeldInventory) GetIterator() InventoryIterator {
	return NewHandHeldInventoryIterator(h.LeftHandItem, h.RightHandItem)
}

type InventoryIterator interface {
	// CQRS is followed. Command and query responsibilities are segregated.
	HasNext() bool
	Next()
	GetCurrent() int
}

type HandHeldInventoryIterator struct {
	//HandHeldInventory HandHeldInventory
	// we cannot have circular dependency in Golang, so, add all var of that struct.

	LeftHandItem  int
	RightHandItem int
}

func NewHandHeldInventoryIterator(left, right int) InventoryIterator {
	return HandHeldInventoryIterator{
		LeftHandItem:  left,
		RightHandItem: right,
	}
}
func (h HandHeldInventoryIterator) HasNext() bool {
	if h.LeftHandItem == 0 && h.RightHandItem == 0 {
		return false
	}
	return true
}
func (h HandHeldInventoryIterator) Next() {
	// do something
}
func (h HandHeldInventoryIterator) GetCurrent() int {
	return 1
}

func main() {
	var inv Inventory
	inv = NewHandHeldInventory(1, 1)
	iter := inv.GetIterator()
	for iter.HasNext() {
		fmt.Println(iter.GetCurrent())
		iter.Next()
	}

	// Code from line 54 to 58 is same, so it can be made as a method and reused for any Inventory (Handheld, bag, warehouse)
}
