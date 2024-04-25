package main

import "fmt"

func main() {
	fmt.Print("Same")
	fmt.Print("line")

	fmt.Println("New")
	fmt.Println("Line")

	x := 3.14
	xs := fmt.Sprint(x)

	fmt.Println("The value of x is:" + xs)
	fmt.Println("The value of x is:", x)

	fmt.Printf("The value of x is: %.2f", x)

	a := 1
	b := 1.1
	c := false
	d := "some_text"

	fmt.Printf("\n%d %f %1.f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
