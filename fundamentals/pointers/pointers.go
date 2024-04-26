package main

import "fmt"

func main() {
	i := 1
	var p *int = nil

	p = &i
	*p++
	fmt.Println(p, *p, i, &i)
}
