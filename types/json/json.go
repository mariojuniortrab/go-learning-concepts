package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := product{1, "Notebook", 1989.90, []string{"sale", "electronics"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	var p2 product
	p2Json := `{"id":2,"name":"pen","price":2.9,"tags":["sale","office"]}`

	json.Unmarshal([]byte(p2Json), &p2)
	fmt.Println(p2)
}
