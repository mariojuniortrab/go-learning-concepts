package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

func main() {
	champion := fastest(
		"https://www.cod3r.com.br",
		"https://www.google.com.br",
		"https://www.youtube.com",
	)

	fmt.Println(champion)

}

func fastest(url1, url2, url3 string) string {

	c1 := titles(url1)
	c2 := titles(url2)
	c3 := titles(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "everybody lost"
		// default:
		// 	return "no response"
	}
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
