package main

import "fmt"

func main() {
	result, _ := factorial(5)
	fmt.Println(result)

	_, err := factorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}

func factorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("invalid number: %d", n)
	case n == 0:
		return 1, nil
	default:
		lastFactorial, _ := factorial(n - 1)
		return n * lastFactorial, nil
	}
}
