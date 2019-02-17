package query

import "github.com/joway/kuafu/document"

type Result struct {
	Documents []document.Document
}

func NewResult(documents []document.Document) *Result {
	return &Result{
		Documents: documents,
	}
}
