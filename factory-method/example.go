package main

import "fmt"

type Vehicle interface {
	GetName() string
}

// inheritance is not possible in Golang
type RoadVehicle interface {
	Vehicle
	GetTiresNum() int
}

type WaterVehicle interface {
	Logistics
	GetWaterCap() int
}

type truck struct{}

func (t truck) GetName() string {
	return "truck"
}
func (t truck) GetTiresNum() int {
	return 6
}

type boat struct{}

func (b boat) GetName() string {
	return "boat"
}

type Logistics interface {
	CreateVehicle() Vehicle
}

type roadLogistics struct{}
func (r roadLogistics) CreateVehicle() Vehicle {
	// returns Road vehicle
	return truck{}
}

type waterLogistics struct {}
func (w waterLogistics) CreateVehicle() Vehicle {
	// returns Water vehicle
	return boat{}
}

func main() {
	r := roadLogistics{}
	fmt.Println(r.CreateVehicle().GetName())
	w := waterLogistics{}
	fmt.Println(w.CreateVehicle().GetName())
}