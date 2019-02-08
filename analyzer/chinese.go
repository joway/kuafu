package analyzer

import (
	"github.com/go-ego/gse"
)

type ChineseAnalyzer struct {
	Analyzer

	seg gse.Segmenter
}

func NewChineseAnalyzer() *ChineseAnalyzer {
	a := &ChineseAnalyzer{}
	if err := a.seg.LoadDict("./data/dict/cn.txt"); err != nil {
		return nil
	}
	return a
}

func (a *ChineseAnalyzer) Tokenize(input string) []string {
	return a.seg.CutSearch(input)
}
