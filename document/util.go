package document

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

func GenerateDocId() string {
	ts := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(ts.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(ts), entropy).String()
}
