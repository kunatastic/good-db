package storage

type Storage interface {
	Get(key string) ([]byte, error)
	Put(key string, value []byte) error
	Delete(key string) error
	Close() error
}

type MemoryStorage struct {
	data map[string][]byte
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string][]byte),
	}
}

func (ms *MemoryStorage) Get(key string) ([]byte, error) {
	value, exists := ms.data[key]
	if !exists {
		return nil, ErrKeyNotFound
	}
	return value, nil
}

func (ms *MemoryStorage) Put(key string, value []byte) error {
	ms.data[key] = value
	return nil
}

func (ms *MemoryStorage) Delete(key string) error {
	delete(ms.data, key)
	return nil
}

func (ms *MemoryStorage) Close() error {
	return nil
}