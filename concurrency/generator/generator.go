package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func main() {
	t1 := titles("https://www.cod3r.com.br", "https://www.google.com.br")
	t2 := titles("https://www.amazon.com", "https://www.youtube.com")

	fmt.Println("firsts:", <-t1, " - ", <-t2)
	fmt.Println("seconds:", <-t1, " - ", <-t2)

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
