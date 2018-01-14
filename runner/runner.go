package runner

import (
	"log"

	"strconv"

	"github.com/tebben/moonfolio/coindata"
	"github.com/tebben/moonfolio/configuration"
	"github.com/tebben/moonfolio/cryptocompare"
	"github.com/tebben/moonfolio/transactions"
)

var (
	allCoins  *cryptocompare.CoinList
	userCoins map[string]*coindata.CoinData
	config    *configuration.Config
)

// Start starts the runner
func Start(c *configuration.Config) {
	config = c
	getCoinList()
	createCoinData()
}

func getCoinList() {
	c, err := cryptocompare.GetCoinList()
	if err != nil {
		// ToDo set error in interface
	}

	allCoins = c
}

func createCoinData() {
	ta, err := transactions.GetTransactions()
	if err != nil {
		log.Fatalf("unable to get transactions: %v", err)
	}

	userCoins = make(map[string]*coindata.CoinData, 0)

	// add transaction to a user coin, create user coin if there is no coin yet
	for _, t := range ta {
		if uc, ok := userCoins[t.CoinID]; !ok {
			uc = &coindata.CoinData{}

			// not created, find coin in allCoins
			if ccCoin, ok2 := allCoins.Data[t.CoinID]; ok2 {
				uc.SetName(ccCoin.FullName)
				uc.SetSymbol(ccCoin.Symbol)
				rank, _ := strconv.Atoi(ccCoin.SortOrder)
				uc.SetRank(rank)
			} else { // coins from cryptocompare did not find the coin refered in the transaction
				uc.SetName(t.CoinID)
				uc.SetSymbol(t.CoinID)
			}

			uc.AddTransaction(t)
			userCoins[uc.Symbol] = uc
		} else {
			uc.AddTransaction(t)
		}
	}
}

func update(price, histoHour, histoDay bool) {

}

// set error in GUI
func setError() {

}
