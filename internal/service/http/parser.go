package http

import "github.com/mrtechit/backend-assignment/internal/service"

type RestService struct {
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
