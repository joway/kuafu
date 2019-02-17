package document

type DocID string

type Document struct {
	Id    DocID
	Index string

	Value string
}
