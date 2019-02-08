package indexer

import (
	"github.com/joway/kuafu/analyzer"
	"github.com/joway/kuafu/document"
	"github.com/joway/kuafu/storage"
)

type Indexer struct {
	storage  storage.Storage
	analysis analyzer.Analyzer
}

func New(storage storage.Storage, analysis analyzer.Analyzer) *Indexer {
	return &Indexer{
		storage:  storage,
		analysis: analysis,
	}
}

func (i *Indexer) Index(doc document.Document) {
	//tokens := i.analyzer.Tokenize(doc.Value)
	if err := i.storage.Set(doc.Id, doc.Value); err != nil {
		return
	}
}
