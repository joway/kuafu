package analyzer

import (
	"fmt"
	"github.com/joway/yaba"
)

type ChineseAnalyzer struct {
	Analyzer

	seg yaba.Segmenter
}

func NewChineseAnalyzer() *ChineseAnalyzer {
	a := &ChineseAnalyzer{}
	_ = a.seg.LoadStopWords("")
	_ = a.seg.LoadDictionary("../data/dict/cn.txt")
	return a
}

func (a *ChineseAnalyzer) Tokenize(input string) []*Token {
	items := a.seg.Segment(input)
	tokens := make([]*Token, 0)
	for _, item := range items {
		fmt.Println(item.Text)
		token := NewToken(
			item.Text,
			[]uint{uint(item.Start)},
		)
		tokens = append(tokens, token)
	}
	return tokens
}
