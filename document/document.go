package document

type Document struct {
	Id    string
	Value string
}

func (d *Document) Data() string {
	return d.Value
}
