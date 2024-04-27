package math

func average(numbers ...float64) float64 {
	total := 0.0

	if len(numbers) == 0 {
		return 0
	}

	for _, num := range numbers {
		total += num
	}

	return total / float64(len(numbers))

}
