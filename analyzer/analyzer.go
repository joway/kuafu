package analyzer

type Analyzer interface {
	Tokenize(input string) []string
}

func New(analysisType string) Analyzer {
	switch analysisType {
	case "chinese":
		return NewChineseAnalyzer()
	default:
		return nil
	}
}
