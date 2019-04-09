package indexer

import (
	"fmt"
	"testing"
)

func TestSegment(t *testing.T) {
	segment := NewSegment()
	segItm1 := NewSegmentItem("1", 3)
	segItm2 := NewSegmentItem("2", 2)
	segItm3 := NewSegmentItem("3", 1)

	segment.Add("a", segItm1)
	segment.Add("a", segItm2)
	segment.Add("b", segItm3)

	ret := segment.Get("a")
	fmt.Println(ret)
}
