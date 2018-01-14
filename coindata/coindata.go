package coindata

import (
	"github.com/tebben/moonfolio/cryptocompare"
	"github.com/tebben/moonfolio/transactions"
)

// CoinData contains information
type CoinData struct {
	Name         string
	Symbol       string
	Rank         int32
	PriceUSD     float64
	HistoHour    *cryptocompare.Histo
	HistoDay     *cryptocompare.Histo
	transactions []transactions.Transaction
}

func (c *CoinData) Update() {

}

func (c *CoinData) SetTransactions() {

}

func (c *CoinData) GetTransactions() {

}

func (c *CoinData) AddTransaction() {

}
