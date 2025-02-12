package db

type Db interface {
	Get(key string) ([][]byte, error)
	Put(key string, value [][]byte) error
}
