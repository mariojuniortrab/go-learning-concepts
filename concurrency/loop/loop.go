package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 30)
	go primes(cap(c), c)

	for prime := range c {
		fmt.Println(prime)
	}

}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func primes(n int, c chan int) {
	first := 2
	for i := 1; i < n; i++ {
		for prime := first; ; prime++ {
			if isPrime(prime) {
				c <- prime
				first = prime + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(c)
}
