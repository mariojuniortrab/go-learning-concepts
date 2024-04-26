package main

import "fmt"

func main() {
	buyProduct1, buyProduct2, buyProduct3 := buy(true, true)

	fmt.Printf("buy product1: %t, buy product 2: %t, buy product3: %t, saving money:%t",
		buyProduct1, buyProduct2, buyProduct3, !(buyProduct1 && buyProduct2 && buyProduct3))
}

func buy(cond1, cond2 bool) (bool, bool, bool) {
	buyProduct1 := cond1 && cond2
	buyProduct2 := cond1 != cond2
	buyProduct3 := cond1 || cond2

	return buyProduct1, buyProduct2, buyProduct3
}
