package main

import "fmt"

func main() {
	fmt.Println(gradeToConcept(9.8))
	fmt.Println(gradeToConcept(10))
	fmt.Println(gradeToConcept(11))
	fmt.Println(gradeToConcept(5))
	fmt.Println(gradeToConcept(3))
}

func gradeToConcept(grade float64) string {
	intGrade := int(grade)

	switch intGrade {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8:
		return "B"
	case 5, 6, 7:
		return "C"
	case 4, 3:
		return "D"
	case 0, 1, 2:
		return "E"
	default:
		return "invalid"
	}

}
