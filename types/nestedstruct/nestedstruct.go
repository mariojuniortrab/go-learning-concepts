package main

import "fmt"

type item struct {
	productID int
	amount    float64
	price     float64
}

type order struct {
	userID int
	items  []item
}

func (o order) totalValue() float64 {
	result := 0.0

	for _, item := range o.items {
		result += item.amount * item.price
	}

	return result
}

func main() {
	order := order{
		userID: 1,
		items: []item{
			{1, 2, 12.10},
			{2, 1, 23.49},
			{11, 100, 3.13},
		},
	}

	fmt.Printf("Total: %.2f", order.totalValue())
}
