package rpcclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Client struct {
	endpoint string
}

// RPCResponse represents a JSON-RPC response
type RPCResponse struct {
	ID     int             `json:"id"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  *RPCError       `json:"error,omitempty"`
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
		fmt.Println("Error:", err)
		return 0, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}

	var rpcResp RPCResponse
	err = json.Unmarshal(body, &rpcResp)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}

	if rpcResp.Error != nil {
		fmt.Printf("Error: %d - %s\n", rpcResp.Error.Code, rpcResp.Error.Message)
		return 0, err
	}

	blockNumber, err := strconv.ParseInt(string(rpcResp.Result[2:]), 16, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	}
	return int(blockNumber), nil
}
