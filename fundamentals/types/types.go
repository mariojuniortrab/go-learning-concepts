package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Integer Literal is :", reflect.TypeOf(32000))

	var b byte = 255
	fmt.Println("Byte is:", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("Max int value is", i1)
	fmt.Println("i1 is:", reflect.TypeOf(i1))

	var i2 rune = 'a'
	fmt.Println("rune is: ", reflect.TypeOf(i2))
	fmt.Println(i2)

	var x float64 = 49.99
	fmt.Println("the type of x is:", reflect.TypeOf(x))
	fmt.Println("Literal value of 49.99 is:", reflect.TypeOf(49.99))

	bo := true
	fmt.Println("type of variable bo is:", reflect.TypeOf(bo))
	fmt.Println(!bo)

	s1 := "Hello, my name is Mario"
	fmt.Println(s1 + "!")
	fmt.Println("This string length is:", len(s1))

	s2 := `Hello
	my
	name is 
	Mario`
	fmt.Println("The size of the other string is: ", len(s2))
}
