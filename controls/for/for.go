package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for 1 <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d", i)
	}

	for {
		fmt.Println("for ever ...")
		time.Sleep(time.Second)
	}
}
