package main

import "fmt"

type grade float64

func (g grade) isBetween(begin, end float64) bool {
	return float64(g) >= begin && float64(g) <= end
}

func gradeForConcept(g grade) string {
	switch {
	case g.isBetween(9.0, 10.0):
		return "A"
	case g.isBetween(8, 9):
		return "B"
	case g.isBetween(5, 8):
		return "C"
	case g.isBetween(3, 5):
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(gradeForConcept(9.8))
	fmt.Println(gradeForConcept(6))
	fmt.Println(gradeForConcept(8.3))
	fmt.Println(gradeForConcept(2))
}
