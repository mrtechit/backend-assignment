package handler

import (
	"github.com/mrtechit/backend-assignment/internal/db"
	"github.com/mrtechit/backend-assignment/internal/db/inmem"
	"github.com/mrtechit/backend-assignment/internal/service"
)

type RestService struct {
	db db.Db
}

func New() *RestService {
	db := inmem.New()
	return &RestService{db: db}
}

func (r RestService) GetCurrentBlock() int {
	//TODO implement me
	panic("implement me")
}

func (r RestService) Subscribe(address string) bool {
	//TODO implement me
	panic("implement me")
}

func (r RestService) GetTransactions(address string) []service.Transaction {
	//TODO implement me
	panic("implement me")
}
