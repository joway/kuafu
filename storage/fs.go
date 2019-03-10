package storage

type FSStorage struct {
	Storage
}

func NewFSStorage() Storage {
	return &FSStorage{}
}

func (s *FSStorage) Get(k string) ([]byte, error) {
	return make([]byte, 0), nil
}

func (s *FSStorage) Add(k string, v []byte) error {
	return nil
}
