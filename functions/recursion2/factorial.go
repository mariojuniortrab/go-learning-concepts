package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

func factorial(n uint) uint {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
