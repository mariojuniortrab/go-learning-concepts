package main

import "fmt"

type car struct {
	name  string
	speed int
}

type ferrari struct {
	car
	isTurboOn bool
}

func main() {
	f := ferrari{}

	f.name = "F40"
	f.speed = 0
	f.isTurboOn = false

	fmt.Printf("Ferrari %s is using turbo? %t", f.name, f.isTurboOn)
}
