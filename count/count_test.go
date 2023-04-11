package count

import (
	"bytes"
	"io"
	"testing"
)

func TestLines(t *testing.T) {
	tests := []struct {
		r    io.Reader
		want int64
	}{
		{&bytes.Buffer{}, 0},
		{bytes.NewBuffer([]byte("abc")), 1},
		{bytes.NewBuffer([]byte("abc\ndef")), 2},
	}
	for i, test := range tests {
		got, err := Lines(test.r)
		if err != nil {
			t.Errorf("Count() failed: %v", err)
		}
		if got != test.want {
			t.Errorf("test %d: got %d, want %d", i, got, test.want)
		}
	}
}
