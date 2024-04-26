package main

import (
	"fmt"
	"strings"
)

type person struct {
	firstname string
	lastname  string
}

func (p person) getFullName() string {
	return p.firstname + " " + p.lastname
}

func (p *person) setFullName(fullname string) {
	slices := strings.Split(fullname, " ")
	p.firstname = slices[0]
	p.lastname = slices[1]
}

func main() {
	p1 := person{"Mario", "Junior"}
	fmt.Println(p1.getFullName())

	p1.setFullName("Airton Senna")
	fmt.Println(p1.getFullName())
}
