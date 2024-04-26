package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 4, 56, 7}

	for i, number := range numbers {
		fmt.Printf("%d) %d\n", i, number)
	}

	for _, number := range numbers {
		fmt.Println(number)
	}

}
