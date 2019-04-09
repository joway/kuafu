package storage

import "errors"

type MemStorage struct {
	Storage

	store map[string][]byte
}

func NewMemStorage() Storage {
	return &MemStorage{
		store: map[string][]byte{},
	}
}

func (s MemStorage) Get(k string) ([]byte, error) {
	if s.store[k] == nil {
		return nil, errors.New("key not found")
	}
	return s.store[k], nil
}

func (s *MemStorage) Add(k string, v []byte) error {
	s.store[k] = v
	return nil
}
