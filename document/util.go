package document

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

func GenerateDocId() DocId {
	ts := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(ts.UnixNano())), 0)
	return DocId(ulid.MustNew(ulid.Timestamp(ts), entropy).String())
}
