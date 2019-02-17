package storage

type Storage interface {
	Get(k string) ([]byte, error)
	Add(k string, v []byte) error
}

func New(storageType string) Storage {
	switch storageType {
	case "fs":
		return NewFSStorage()
	case "mem":
		return NewMemStorage()
	default:
		return nil
	}
}
