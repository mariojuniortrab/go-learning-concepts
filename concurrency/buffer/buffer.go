package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	go routine(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}

func routine(ch chan int) {
	fmt.Println("Executed")

	ch <- 1
	fmt.Println("Executed 1")
	ch <- 2
	fmt.Println("Executed 2")
	ch <- 3
	fmt.Println("Executed 3")
	ch <- 4
	fmt.Println("Executed 4")
	ch <- 5
	fmt.Println("Executed 5")
	ch <- 6
}
