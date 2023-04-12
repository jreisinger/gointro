package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/jreisinger/gointro/byte"
)

var counter byte.Counter

func main() {
	urls := os.Args[1:]
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch) // start a goroutine
	}
	for range urls {
		fmt.Println(<-ch) // receive from channel
	}
	fmt.Printf("----- -------\n")
	fmt.Printf("%5.3f %7d\n", time.Since(start).Seconds(), counter)
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()
	n, err := io.Copy(&counter, resp.Body)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	ch <- fmt.Sprintf("%5.3f %7d %s", time.Since(start).Seconds(), n, url)
}
