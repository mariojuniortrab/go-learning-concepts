package main

import "fmt"

func main() {
	fmt.Println(exec(multiplication, 2, 3 ))
}

func multiplication(a, b int) int {
	return a * b
}

func exec(function func(int, int) int, p1, p2 int) int {
	return function(p1, p2)
}
