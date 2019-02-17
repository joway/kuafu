package main

import (
	"fmt"
	"github.com/joway/kuafu/analyzer"
	"github.com/joway/kuafu/document"
	"github.com/joway/kuafu/indexer"
	"github.com/joway/kuafu/query"
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
	//prepare document
	doc := document.Document{Id: "1", Value: "hi, kuafu"}
	//index document
	idx.Index(doc)
	//setup query
	opt := query.NewOption(1000)
	q := query.NewQuery("myindex", "kuafu", *opt)
	//create searcher
	sch := searcher.New(ana)
	//search document
	ret := sch.Search(q)

	fmt.Println(ret)
}
