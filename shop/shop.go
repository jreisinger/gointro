// Shop is a simple e-shop showing price of shoes and socks.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe(":8000", db))
}

type database map[string]euro

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

type euro float32

func (e euro) String() string {
	return fmt.Sprintf("€%.2f", e)
}
