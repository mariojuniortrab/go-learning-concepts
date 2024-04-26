package main

import "fmt"

func main() {
	var grade [3]float64

	grade[0], grade[1], grade[2] = 0.1, 5.2, 7

	fmt.Println(grade)

	total := 0.0

	for i := 0; i < len(grade); i++ {
		total += grade[i]
	}

	average := total / float64(len(grade))

	fmt.Printf("Average:%.2f\n", average)
}
