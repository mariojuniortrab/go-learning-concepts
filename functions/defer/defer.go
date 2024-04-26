package main

import "fmt"

func main() {
	fmt.Println(getApprovedValue(6000))
	fmt.Println(getApprovedValue(3000))
}

func getApprovedValue(number int) int {
	defer fmt.Println("End!")

	if number > 5000 {
		fmt.Println("High value...")
	}

	fmt.Println("low value...")
	return number
}
