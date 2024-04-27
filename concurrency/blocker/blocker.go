package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go routine(c)
	fmt.Println(<-c)
	fmt.Println("read")
	fmt.Println(<-c)
	fmt.Println("end")
}

func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 1
	fmt.Println("After being read")
}
