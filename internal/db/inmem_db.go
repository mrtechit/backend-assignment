package db

import "sync"

type Database struct {
	db   map[string][][]byte
	lock sync.RWMutex
}

func (d Database) Get(key string) ([][]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (d Database) Put(key string, value [][]byte) error {
	//TODO implement me
	panic("implement me")
}
