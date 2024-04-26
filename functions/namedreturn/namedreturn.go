package main

import "fmt"

func main() {
	x, y := change(3, 4)
	fmt.Println(x, y)

	second, first := change(1, 2)
	fmt.Println(second, first)
}

func change(p1, p2 int) (second int, first int) {
	second = p1
	first = p2

	return
}
