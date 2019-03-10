package analyzer

import (
	"fmt"
	"testing"
)

func TestChineseAnalyzer_Tokenize(t *testing.T) {
	a := NewChineseAnalyzer()
	tokens := a.Tokenize("今天的天气真的很好啊，但是明天天气可能会很差")
	for _, token := range tokens {
		fmt.Println(token.Value)
	}

	tokens = a.Tokenize("Hi, how are you?")
	for _, token := range tokens {
		fmt.Println(token.Value)
	}
}
