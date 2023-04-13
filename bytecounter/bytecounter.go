package main

import (
	"fmt"
)

// User-defined type (based on int) that satisfies the io.Writer interface.
type ByteCounter int

// Use a pointer receiver (bc) if you want to change it, otherwise it's a copy.
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
