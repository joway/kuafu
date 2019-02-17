package storage

type MemStorage struct {
	Storage
}

func NewMemStorage() *MemStorage {
	return &MemStorage{}
}

func (s *MemStorage) Get(k string) ([]byte, error) {
	return make([]byte, 0), nil
}

func (s *MemStorage) Add(k string, v []byte) error {
	return nil
}
