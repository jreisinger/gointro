// Package fetch works with HTTP resources.
package fetch

import (
	"fmt"
	"io"
	"net/http"
)

// Fetch gets the resource at url and sends its size in bytes on ch.
func Fetch(url string, ch chan string) {
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
