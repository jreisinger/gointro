package byte

import "sync"

var mu sync.Mutex

type Counter int

func (c *Counter) Write(p []byte) (int, error) {
	mu.Lock()
	*c += Counter(len(p))
	mu.Unlock()
	return len(p), nil
}
