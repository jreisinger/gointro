package main

import (
	"fmt"
	"os"

	"github.com/jreisinger/gointro/fetch"
)

func main() {
	urls := os.Args[1:]
	ch := make(chan string)
	for _, url := range urls {
		go fetch.Fetch(url, ch) // start a goroutine
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel
	}
}
