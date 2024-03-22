package inmem

import (
	"errors"
	"sync"
)

type Database struct {
	db   map[string][][]byte
	lock sync.RWMutex
}

func New() *Database {
	return &Database{
		db: make(map[string][][]byte),
	}
}

func (d Database) Get(key string) ([][]byte, error) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	if entry, ok := d.db[key]; ok {
		result := make([][]byte, len(entry))
		copy(result, entry)
		return result, nil
	}
	return nil, errors.New("not found")
}

func (d Database) Put(key string, value [][]byte) error {
	d.lock.RLock()
	defer d.lock.RUnlock()
	d.db[key] = append(d.db[key], value...)
	return nil
}
