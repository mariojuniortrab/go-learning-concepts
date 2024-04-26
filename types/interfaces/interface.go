package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	firstname string
	lastname  string
}

type product struct {
	name  string
	price float64
}

func (p person) toString() string {
	return p.firstname + " " + p.lastname
}

func (p product) toString() string {
	return fmt.Sprintf("%s: US$%.2f", p.name, p.price)
}

func print(p printable) {
	fmt.Println(p.toString())
}

func main() {
	var thing printable = person{"Mario", "Junior"}
	print(thing)

	thing = product{"t-shirt", 79.90}
	print(thing)

	thing2 := product{"jeans", 80.00}
	print(thing2)
}
