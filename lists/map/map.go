package main

import "fmt"

func main() {
	//var approved map[int]string
	approved := make(map[int]string)

	approved[121212] = "mary"
	approved[121211] = "joseph"
	approved[121221] = "ann"

	fmt.Println(approved)

	for id, name := range approved {
		fmt.Printf("%s (id: %d)\n", name, id)
	}

	fmt.Println(approved[121212])
	delete(approved, 121212)

	fmt.Println(approved)

}
