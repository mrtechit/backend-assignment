package handler

import (
	"github.com/mrtechit/backend-assignment/internal/db"
	"github.com/mrtechit/backend-assignment/internal/db/inmem"
	"github.com/mrtechit/backend-assignment/internal/service"
	"github.com/mrtechit/backend-assignment/pkg/rpcclient"
)

type RestService struct {
	db        db.Db
	rpcClient *rpcclient.Client
}

func New() *RestService {
	db := inmem.New()
	rpcClient := rpcclient.New("https://cloudflare-eth.com")
	return &RestService{db: db,
		rpcClient: rpcClient}
}

func (r RestService) GetCurrentBlock() (int, error) {
	res, err := r.rpcClient.BlockNumber()
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (r RestService) Subscribe(address string) bool {
	//TODO implement me
	panic("implement me")
}

func (r RestService) GetTransactions(address string) []service.Transaction {
	//TODO implement me
	panic("implement me")
}
