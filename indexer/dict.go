package indexer

import "github.com/joway/kuafu/document"

type DictItem struct {
	DocId document.DocID
	TF  uint
}

type Dict struct {
	store map[string][]DictItem
}

func NewDict() *Dict {
	store := make(map[string][]DictItem, 0)
	return &Dict{store}
}

func (d *Dict) Get(token string) []DictItem {
	return d.store[token]
}

func (d *Dict) Add(token string, item DictItem) {
	items := d.store[token]
	d.store[token] = append(items, item)
}
