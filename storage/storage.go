package storage

type Storage interface {
	Get(k string) (string, error)
	Set(k string, v string) error
}

func New(storageType string) Storage {
	switch storageType {
	case "local":
		return NewLocalStorage()
	default:
		return nil
	}
}
