package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2
	b := 3

	fmt.Println("sum => ", a+b)
	fmt.Println("subtract => ", a-b)
	fmt.Println("division => ", a/b)
	fmt.Println("multiplication => ", a*b)
	fmt.Println("mod => ", a%b)

	//bitwise
	fmt.Println("and => ", a&b)
	fmt.Println("or => ", a|b)
	fmt.Println("xor => ", a^b)

	c, d := 3.0, 2.0

	fmt.Println("max => ", math.Max(float64(a), float64(b)))
	fmt.Println("min => ", math.Min(c, d))
	fmt.Println("pow => ", math.Pow(c, d))
}
