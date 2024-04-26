package main

import "fmt"

func main() {
	employersByWord := map[string]map[string]float64{
		"m": {
			"mario junior":  15000.00,
			"michael james": 12300.23,
		},
		"b": {
			"brittany swords": 16980.00,
		},
		"g": {
			"gerard muller": 11234.3,
		},
		"p": {
			"peter wayne": 14000.45,
		},
	}

	delete(employersByWord, "p")

	fmt.Println(employersByWord)
}
