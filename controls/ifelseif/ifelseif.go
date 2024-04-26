package main

import "fmt"

func main() {
	fmt.Println(gradeToConcept(9.1))
	fmt.Println(gradeToConcept(8.1))
	fmt.Println(gradeToConcept(7.1))
	fmt.Println(gradeToConcept(4.1))
	fmt.Println(gradeToConcept(2.1))
}

func gradeToConcept(grade float64) string {
	if grade >= 9 && grade <= 10 {
		return "A"
	} else if grade >= 8 {
		return "B"
	} else if grade >= 5 {
		return "C"
	} else if grade >= 3 {
		return "D"
	}

	return "E"
}
