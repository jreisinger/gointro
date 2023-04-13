package main

import (
	"fmt"
)

type ByteCounter int

func (bc *ByteCounter) Write(p []byte) (int, error) {
	*bc += ByteCounter(len(p)) // must explicitly convert int to ByteCounter
	return len(p), nil
}

func main() {
	s := "hello"
	var bc ByteCounter
	fmt.Fprint(&bc, s)
	fmt.Println(bc)
}
