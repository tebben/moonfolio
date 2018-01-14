package runner

import (
	"log"

	"github.com/tebben/moonfolio/configuration"
	"github.com/tebben/moonfolio/cryptocompare"
	"github.com/tebben/moonfolio/transactions"
)

var (
	coinList        *cryptocompare.CoinList
	transactionList []*transactions.Transaction
)

// Start starts the runner
func Start(config *configuration.Config) {
	// get coin list

	// get transaction
}

func getCoinList() {
	c, err := cryptocompare.GetCoinList()
	if err != nil {

	}

	coinList = c
}

func getTransactions() {
	t, err := transactions.GetTransactions()
	if err != nil {
		log.Fatalf("unable to get transactions: %v", err)
	}

	transactionList = t
}

// set error in GUI
func setError() {

}
