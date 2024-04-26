package main

import "fmt"

type course struct {
	name string
}

func main() {
	var thing interface{}
	fmt.Println(thing)

	thing = 3

	type dynamic interface{}
	var thing2 dynamic = "hi"
	fmt.Println(thing2)

	thing2 = true
	fmt.Println(thing2)

	thing2 = course{"course 1"}
	fmt.Println(thing2)

}
