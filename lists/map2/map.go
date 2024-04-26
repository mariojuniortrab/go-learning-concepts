package main

import "fmt"

func main() {
	employersWage := map[string]float64{
		"mario junior":    15000.0,
		"john smith":      14234.2,
		"brittany swords": 16123.34,
	}

	employersWage["gerard muller"] = 11000.29

	delete(employersWage, "michael james")

	for name, wage := range employersWage {
		fmt.Printf("%s - %.2f\n", name, wage)
	}
}
