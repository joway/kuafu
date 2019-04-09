package searcher

import (
	"github.com/joway/kuafu/document"
	"github.com/joway/kuafu/query"
)

type Searcher struct {
}

func New() Searcher {
	return Searcher{}
}

func (s Searcher) Search(q query.Query) query.Result {
	//tokens := s.analyzer.Tokenize(q.Body)
	var docs []document.Document
	result := query.NewResult(docs)
	return result
}
