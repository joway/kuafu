package analyzer

import (
	"github.com/wangbin/jiebago"
)

type ChineseAnalyzer struct {
	Analyzer

	seg jiebago.Segmenter
}

func NewChineseAnalyzer() Analyzer {
	a := &ChineseAnalyzer{}
	if err := a.seg.LoadDictionary("./data/dict/cn.txt"); err != nil {
		return nil
	}
	return a
}

func (a *ChineseAnalyzer) Tokenize(input []byte) []Token {
	ch := a.seg.CutForSearch(string(input), true)

	tokens := make([]Token, 0)
	for word := range ch {
		token := NewToken(
			word,
			[]uint{0},
		)
		tokens = append(tokens, token)
	}
	return tokens
}
