package main

import "fmt"

func main() {
	printResult(7.2)
	printResult(5)
}

func printResult(grade float64) {
	if grade >= 7 {
		fmt.Println("Success")
	} else {
		fmt.Println("Failed")
	}
}
