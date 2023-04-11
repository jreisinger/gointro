package count

import (
	"bufio"
	"io"
)

func Bytes(r io.Reader) (int64, error) {
	return io.Copy(io.Discard, r)
}

func Lines(r io.Reader) (int64, error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	var nLines int
	for s.Scan() {
		nLines++
	}
	return int64(nLines), s.Err()
}
