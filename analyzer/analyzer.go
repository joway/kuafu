package analyzer

type Token struct {
	Value   string `json:"value"`
	TF      uint   `json:"tf"`
	offsets []uint
}

type Analyzer interface {
	Tokenize(input []byte) []Token
}

func NewToken(value string, offsets []uint) Token {
	return Token{
		Value:   value,
		TF:      uint(len(offsets)),
		offsets: offsets,
	}
}
