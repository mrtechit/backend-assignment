package db

type db interface {
	Get(key string) ([][]byte, error)
	Put(key string, value [][]byte) error
}