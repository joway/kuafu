package searcher

import "github.com/joway/kuafu/query"

type Searcher struct {
}

func New() *Searcher {
	return &Searcher{}
}

func (s *Searcher) Search(query string) *query.Result {
	return nil
}
