package query

type Option struct {
	Timeout uint //ms
}

type Query struct {
	Index  string
	Body   string
	Option Option
}

func NewOption(timeout uint) *Option {
	return &Option{
		Timeout: timeout,
	}
}

func NewQuery(index string, body string, option Option) *Query {
	return &Query{
		Index:  index,
		Body:   body,
		Option: option,
	}
}
