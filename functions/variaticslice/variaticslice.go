package main

import "fmt"

func main() {
	approved := []string{"mario", "michael", "brittany", "gerard"}
	printApproved(approved...)
}

func printApproved(approved ...string) {
	fmt.Println("Approved list:")
	for i, name := range approved {
		fmt.Printf("%d) %s\n", i+1, name)
	}
}
