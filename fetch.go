package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jreisinger/gosec/count"
)

func main() {
	urls := os.Args[1:]
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%5.3f elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()
	n, err := count.Bytes(resp.Body)
	if err != nil {
		log.Print(err)
	}
	ch <- fmt.Sprintf("%5.3f %5d %s", time.Since(start).Seconds(), n, url)
}
