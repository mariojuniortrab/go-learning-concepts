package main

import "fmt"

func main() {
	fmt.Println(average(5, 4, 5, 6))
}

func average(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}

	return total / float64(len(numbers))
}
