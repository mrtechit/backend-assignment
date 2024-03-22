package indexer

import (
	"github.com/mrtechit/backend-assignment/internal/service"
	"time"
)

type Result struct {
	Number       string                `json:"number"`
	Hash         string                `json:"hash"`
	Transactions []service.Transaction `json:"transactions"`
}

// EthIndexer This is not for production , just a prototype!
func EthIndexer(interval time.Duration) {

	ticker := time.NewTicker(5 * time.Minute)

	for {
		select {
		case <-ticker.C:
			ethIndex()
		}
	}
}

//TODO Logic goes here which compares trx in memory db with new data for subscribed addresses
func ethIndex() {
	//client := rpcclient.New("https://cloudflare-eth.com")
	//head, _ := client.BlockNumber()
	//block, _ := client.BlockByNumber(head)

}
