package main

import (
	"github.com/joway/kuafu/analyzer"
	"github.com/joway/kuafu/document"
	"github.com/joway/kuafu/indexer"
	"github.com/joway/kuafu/searcher"
	"github.com/joway/kuafu/storage"
)

func main() {
	//setup storage
	stg := storage.New("local")
	//setup analyzer
	ana := analyzer.New("chinese")
	//create indexer
	idx := indexer.New(stg, ana)
	//create search
	sch := searcher.New()
	//prepare document
	doc := document.Document{Id: "1", Value: "hi, kuafu"}
	//index document
	idx.Index(doc)
	//search document
	query := ""
	sch.Search(query)
}
