package searcher

import (
	"fmt"
	"github.com/joway/kuafu/analyzer"
	"github.com/joway/kuafu/document"
	"github.com/joway/kuafu/query"
)

type Searcher struct {
	analyzer analyzer.Analyzer
}

func New(analyzer analyzer.Analyzer) *Searcher {
	return &Searcher{analyzer: analyzer}
}

func (s *Searcher) Search(q *query.Query) *query.Result {
	tokens := s.analyzer.Tokenize(q.Body)
	fmt.Println(tokens)

	docs := make([]document.Document, 0)
	result := query.NewResult(docs)
	return result
}
