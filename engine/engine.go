package engine

import "github.com/joway/kuafu/indexer"

type Engine struct {
	indexerStore map[string]indexer.Indexer
}
