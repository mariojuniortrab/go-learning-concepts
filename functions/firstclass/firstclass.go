package main

import "fmt"

func main() {
	fmt.Println(sum(2, 3))

	var sub = func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(5, 3))
}

var sum = func(a, b int) int {
	return a + b
}
