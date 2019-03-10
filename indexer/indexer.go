package indexer

import (
	"github.com/joway/kuafu/analyzer"
	"github.com/joway/kuafu/document"
	"github.com/joway/kuafu/storage"
)

type Indexer struct {
	storage  storage.Storage
	analyzer analyzer.Analyzer

	hotSegment   Segment
	coldSegments []Segment
}

func New(storage storage.Storage, analyzer analyzer.Analyzer) *Indexer {
	return &Indexer{
		storage:  storage,
		analyzer: analyzer,

		hotSegment:   NewSegment(),
		coldSegments: []Segment{},
	}
}

func (i *Indexer) Index(docs []document.Document) {
	for _, doc := range docs {
		tokens := i.analyzer.Tokenize(doc.Data)
		//store raw doc
		if err := i.storage.Add(string(doc.Id), []byte(doc.Data)); err != nil {
			return
		}
		//store reverse index
		for _, token := range tokens {
			i.hotSegment.Add(string(doc.Id), NewSegmentItem(
				doc.Id,
				token.TF,
			))
		}
	}
}

func (i Indexer) Segments() []Segment {
	return append(i.coldSegments, i.hotSegment)
}
