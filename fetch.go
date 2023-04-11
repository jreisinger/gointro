package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jreisinger/gosec/stopwatch"
)

func main() {
	urls := os.Args[1:]
	sw := stopwatch.Start()
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%5.3f elapsed\n", sw.Stop().Seconds())
}

func fetch(url string, ch chan string) {
	sw := stopwatch.Start()
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()
	n, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		log.Print(err)
	}
	ch <- fmt.Sprintf("%5.3f %5d %s", sw.Stop().Seconds(), n, url)
}
