package analyzer

type Token struct {
	Value   string
	TF      uint
	offsets []uint
}

type Analyzer interface {
	Tokenize(input string) []*Token
}

func NewToken(value string, offsets []uint) *Token {
	return &Token{
		Value:   value,
		TF:      uint(len(offsets)),
		offsets: offsets,
	}
}

func New(analysisType string) Analyzer {
	switch analysisType {
	case "chinese":
		return NewChineseAnalyzer()
	default:
		return nil
	}
}
