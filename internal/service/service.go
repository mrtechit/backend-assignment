package service

import (
	"encoding/json"
	"github.com/mrtechit/backend-assignment/internal/service/handler"
	"math/big"
	"net/http"
)

type Parser interface {
	GetCurrentBlock() (int, error)
	Subscribe(address string) bool
	GetTransactions(address string) []Transaction
}

type ErrorResponse struct {
	ErrorMessage string `json:"error"`
}

type Response struct {
	BlockNumber int `json:"block_number"`
}

type Transaction struct {
	ChainID     *big.Int `json:"chainId"`
	BlockNumber *big.Int `json:"blockNumber"`
	Hash        string   `json:"hash"`
	Nonce       *big.Int `json:"nonce"`
	From        string   `json:"from"`
	To          string   `json:"to"`
	Value       *big.Int `json:"value"`
	Gas         *big.Int `json:"gas"`
	GasPrice    *big.Int `json:"gasPrice"`
	Input       string   `json:"input"`
}

func GetCurrentBlock(w http.ResponseWriter, r *http.Request) {
	rs := handler.New()
	blockNumber, err := rs.GetCurrentBlock()
	if err != nil {
		http.Error(w, "Error occured", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := Response{BlockNumber: blockNumber}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func Subscribe(w http.ResponseWriter, r *http.Request) {

}

func GetTransactions(w http.ResponseWriter, r *http.Request) {

}
