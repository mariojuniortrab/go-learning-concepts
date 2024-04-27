package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

func main() {
	c := join(speak("john"), speak("mary"))
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}

func titles(url ...string) <-chan string {
	c := make(chan string)
	for _, url := range url {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := io.ReadAll(resp.Body)
			r := regexp.MustCompile(`<title.*?>(.*)</title>`)
			c <- r.FindAllStringSubmatch(string(html), -1)[0][1]
		}(url)
	}

	return c
}

func speak(person string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando: %d ", person, i)
		}
	}()

	return c
}

func join(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()

	return c
}
