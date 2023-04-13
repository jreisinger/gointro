package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	urls := os.Args[1:]
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch) // start a goroutine
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel
	}
}

func fetch(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()

	nbytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	ch <- fmt.Sprintf("%-7d %s", nbytes, url)
}
