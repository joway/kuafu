package storage

type LocalStorage struct {
	Storage
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{}
}

func (s *LocalStorage) Get(k string) (string, error) {
	return "", nil
}

func (s *LocalStorage) Set(k string, v string) error {
	return nil
}
