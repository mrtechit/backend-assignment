package handler

import (
	"encoding/json"
	"github.com/mrtechit/backend-assignment/internal/db"
	"github.com/mrtechit/backend-assignment/internal/db/inmem"
	"github.com/mrtechit/backend-assignment/internal/service"
	"github.com/mrtechit/backend-assignment/pkg/rpcclient"
	"strings"
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
	if err := r.db.Put(strings.ToLower(address), [][]byte{}); err != nil {
		return false
	}
	return true
}

func (r RestService) GetTransactions(address string) []service.Transaction {
	txs, err := r.db.Get(strings.ToLower(address))
	if err != nil {
		return nil
	}
	transactions := make([]service.Transaction, len(txs))
	for i, j := range txs {
		var tx service.Transaction
		if err := json.Unmarshal(j, &tx); err != nil {
			//TODO handle errors
		}
		transactions[i] = tx
	}
	return transactions
}
