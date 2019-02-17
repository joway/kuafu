package indexer

import (
	"fmt"
	"github.com/joway/kuafu/analyzer"
	"github.com/joway/kuafu/document"
	"github.com/joway/kuafu/storage"
)

type Indexer struct {
	storage  storage.Storage
	analyzer analyzer.Analyzer
}

func New(storage storage.Storage, analyzer analyzer.Analyzer) *Indexer {
	return &Indexer{
		storage:  storage,
		analyzer: analyzer,
	}
}

func (i *Indexer) Index(doc document.Document) {
	tokens := i.analyzer.Tokenize(doc.Value)
	fmt.Println(tokens)
	//if err := i.storage.Add(doc.Id, doc.Value); err != nil {
	//	return
	//}
}
