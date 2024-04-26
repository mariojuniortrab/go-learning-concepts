package main

import "fmt"

type sportCar interface {
	turnTurboOn()
}

type luxuriusCar interface {
	turnAutoPilotOn()
}

type luxuriusSportCar interface {
	sportCar
	luxuriusCar
}

type bmw7 struct{}

func (b *bmw7) turnTurboOn() {
	fmt.Println("turbo is on")
}

func (b *bmw7) turnAutoPilotOn() {
	fmt.Println("The car is going by it's own")
}

func main() {
	var b luxuriusSportCar = &bmw7{}
	b.turnAutoPilotOn()
	b.turnTurboOn()
}
