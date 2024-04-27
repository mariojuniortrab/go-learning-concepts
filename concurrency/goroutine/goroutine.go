package main

import (
	"fmt"
	"time"
)

func main() {
	speak("mary", "Why don't you speak with me?", 3)
	speak("john", "I can only speak just after you!", 2)

	go speak("mary", "hello...", 500)
	go speak("john", "hi...", 500)

}

func speak(person, text string, amount int) {
	for i := 0; i < amount; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteration %d)\n", person, text, i+1)
	}
}
