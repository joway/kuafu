package indexer

import (
	"fmt"
	"github.com/joway/kuafu/document"
	"sort"
)

type SegmentItem struct {
	DocId document.DocId
	TF    uint
}

func NewSegmentItem(docId document.DocId, tf uint) SegmentItem {
	return SegmentItem{
		DocId: docId,
		TF:    tf,
	}
}

func (item SegmentItem) String() string {
	return fmt.Sprintf("<DocId: %s, TF: %v>", item.DocId, item.DocId)
}

type SegmentItemList []SegmentItem

func (l SegmentItemList) Len() int           { return len(l) }
func (l SegmentItemList) Less(i, j int) bool { return l[i].TF < l[j].TF }
func (l SegmentItemList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l SegmentItemList) IsEmpty() bool      { return l.Len() == 0 }
func (l SegmentItemList) First() SegmentItem { return l[0] }
func (l SegmentItemList) Last() SegmentItem  { return l[l.Len()-1] }

type Segment struct {
	store map[string]SegmentItemList
}

func NewSegment() Segment {
	return Segment{
		store: map[string]SegmentItemList{},
	}
}

func (s *Segment) Add(token string, item SegmentItem) {
	if s.store[token] == nil {
		s.store[token] = SegmentItemList{}
	}
	s.store[token] = append(s.store[token], item)
	//TODO: performance issue
	sort.Stable(s.store[token])
}

func (s Segment) Get(token string) SegmentItemList {
	got := s.store[token]
	if got == nil {
		return SegmentItemList{}
	}
	return got
}
