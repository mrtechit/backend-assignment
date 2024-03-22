package rpcclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

type Client struct {
	endpoint string
}

// RPCResponse represents a JSON-RPC response
type RPCResponse struct {
	ID     int       `json:"id"`
	Result string    `json:"result,omitempty"`
	Error  *RPCError `json:"error,omitempty"`
}

type Res struct {
	Number       string        `json:"number"`
	Hash         string        `json:"hash"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	ChainID          string            `json:"chainId"`
	BlockNumber      string            `json:"blockNumber"`
	BlockHash        string            `json:"-"`
	Hash             string            `json:"hash"`
	Nonce            string            `json:"nonce"`
	From             string            `json:"from"`
	To               string            `json:"to"`
	Value            string            `json:"value"`
	Gas              string            `json:"gas"`
	GasPrice         string            `json:"gasPrice"`
	Input            string            `json:"input"`
	Type             string            `json:"-"`
	R                string            `json:"-"`
	S                string            `json:"-"`
	V                string            `json:"-"`
	TransactionIndex string            `json:"-"`
	AccessList       []AccessListEntry `json:"-"`
}

type AccessListEntry struct {
	Address     string   `json:"address"`
	StorageKeys []string `json:"storageKeys"`
}

// RPCError represents an error returned in a JSON-RPC response
type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func New(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

func (c Client) BlockNumber() (int, error) {
	url := c.endpoint

	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_blockNumber",
		"params":  []interface{}{},
		"id":      1,
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return 0, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var rpcResp RPCResponse
	err = json.Unmarshal(body, &rpcResp)
	if err != nil {
		return 0, err
	}

	if rpcResp.Error != nil {
		return 0, err
	}

	blockNumber, err := strconv.ParseInt(string(rpcResp.Result[2:]), 16, 64)
	if err != nil {
		return 0, err
	}
	return int(blockNumber), nil
}

func (c Client) BlockByNumber(blockNumber int) (*RPCResponse, error) {
	url := c.endpoint

	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getBlockByNumber",
		"params":  []string{fmt.Sprintf("0x%x", blockNumber)},
		"id":      rand.Int(),
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rpcResp RPCResponse
	err = json.Unmarshal(body, &rpcResp)
	if err != nil {
		return nil, err
	}

	if rpcResp.Error != nil {
		return nil, err
	}

	return &rpcResp, nil
}
