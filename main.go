package main

import (
	"github.com/joway/kuafu/analyzer"
	"github.com/joway/kuafu/document"
	"github.com/joway/kuafu/indexer"
	"github.com/joway/kuafu/storage"
)

func main() {
	//setup storage
	stg := storage.NewMemStorage()
	//setup analyzer
	ana := analyzer.NewChineseAnalyzer()
	//create indexer
	idx := indexer.New(stg, ana)

	indexName := "myIndex"
	//prepare document
	docId := "1"
	doc := document.New(indexName, document.DocId(docId), []byte("hi, kuafu! how are you?"))
	//index document
	idx.Index([]document.Document{doc})
	//
	////setup query
	//opt := query.NewOption(1000)
	//q := query.NewQuery(indexName, "kuafu", *opt)
	////create searcher
	//sch := searcher.New(ana)
	////search document
	//ret := sch.Search(q)
	//
	//fmt.Println(ret)
}
