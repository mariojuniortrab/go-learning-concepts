package main

import "fmt"

func main() {
	resultado := sum(3, 4)
	print(resultado)
}

func sum(a int, b int) int {
	return a + b
}

func print(a int) {
	fmt.Println(a)
}
