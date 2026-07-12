package id


import (
	"crypto/rand"
	"fmt"
	"sync"
	"time"

	"github.com/oklog/ulid/v2"
)

var (
	mu      sync.Mutex
	entropy = ulid.Monotonic(rand.Reader, 0)
)

func New() string {
	mu.Lock()
	defer mu.Unlock()

	id := ulid.MustNew(
		ulid.Timestamp(time.Now()),
		entropy,
	)

	return fmt.Sprintf("hawk_%s", id.String())
}