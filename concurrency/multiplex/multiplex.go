package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func main() {
	c := join(
		titles("https://www.cod3r.com.br", "https://www.google.com.br"),
		titles("https://www.amazon.com", "https://www.youtube.com"),
	)

	fmt.Println(<-c, " - ", <-c)
	fmt.Println(<-c, " - ", <-c)
}

func redirect(origin <-chan string, destiny chan string) {
	for {
		destiny <- <-origin
	}
}

func join(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go redirect(input1, c)
	go redirect(input2, c)
	return c
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
