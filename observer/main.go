package main

import "fmt"

type Observable interface {
	Add(observer Observer)
	Remove(observer Observer)
	Notify()

	GetTemperature() float32
	SetTemperature(temperature float32)
}

type WeatherStation struct {
	temperature float32
	observers   []Observer
}

func (ws *WeatherStation) Add(observer Observer) {
	fmt.Println("Adding observer", observer.GetID())
	ws.observers = append(ws.observers, observer)
}

func (ws *WeatherStation) Remove(observer Observer) {
	fmt.Println("Removing", observer.GetID())
	for index, o := range ws.observers {
		if o.GetID() == observer.GetID() {
			ws.observers = append(ws.observers[:index], ws.observers[index+1:]...)
			break
		}
	}
}

func (ws *WeatherStation) Notify() {
	for _, observer := range ws.observers {
		observer.Update()
	}
}

func (ws *WeatherStation) GetTemperature() float32 {
	return ws.temperature
}

func (ws *WeatherStation) SetTemperature(temperature float32) {
	ws.temperature = temperature
}

type Observer interface {
	Update()
	GetID() string
}

type Display struct {
	id      string
	station WeatherStation
}

func (d Display) Update() {
	fmt.Println("Updating", d.id, "display with temperature", d.station.GetTemperature())
}

func (d Display) GetID() string {
	return d.id
}

func main() {
	station := &WeatherStation{
		temperature: 1.0,
		observers:   []Observer{},
	}
	phoneDisplay := Display{
		id:      "phone",
		station: *station,
	}
	station.Add(phoneDisplay)
	windowDisplay := Display{
		id:      "window",
		station: *station,
	}
	station.Add(windowDisplay)
	station.Notify()

	station.Remove(windowDisplay)
	station.Notify()

	station.Add(windowDisplay)
	station.Notify()
}

