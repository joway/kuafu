package document

type DocId string

type Document struct {
	IndexName string
	Id        DocId
	Data      []byte
}

func New(index string, id DocId, data []byte) Document {
	return Document{
		IndexName: index,
		Id:        id,
		Data:      data,
	}
}
