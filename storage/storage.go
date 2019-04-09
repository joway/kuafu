package storage

type Storage interface {
	Get(k string) ([]byte, error)
	Add(k string, v []byte) error
}
