package main

import "fmt"

type sportCar interface {
	turnTurboOn()
}

type ferrari struct {
	model     string
	isTurboOn bool
	speed     float64
}

func (f *ferrari) turnTurboOn() {
	f.isTurboOn = true
	f.speed += 10
}

func main() {
	car1 := ferrari{"F40", false, 0}
	car1.turnTurboOn()

	var car2 sportCar = &ferrari{"F430", false, 0}
	car2.turnTurboOn()

	fmt.Println(car1, car2)
}
