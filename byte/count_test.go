package byte

import (
	"testing"
)

func TestCounter(t *testing.T) {
	tests := []struct {
		p    []byte
		want int
	}{
		{[]byte{}, 0},
		{[]byte("abc"), 3},
		{[]byte("abc\ndef"), 7},
	}
	var c Counter
	for i, test := range tests {
		got, err := c.Write(test.p)
		if err != nil {
			t.Errorf("Count() failed: %v", err)
		}
		if got != test.want {
			t.Errorf("test %d: got %d, want %d", i, got, test.want)
		}
	}
	if c != 10 {
		t.Errorf("go total count of %d, want %d", c, 10)
	}
}
