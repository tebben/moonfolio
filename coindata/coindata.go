package coindata

import (
	"github.com/tebben/moonfolio/cryptocompare"
	"github.com/tebben/moonfolio/transactions"
)

// CoinData contains information
type CoinData struct {
	Name         string
	Symbol       string
	Rank         int
	PriceUSD     float64
	HistoHour    *cryptocompare.Histo
	HistoDay     *cryptocompare.Histo
	transactions []*transactions.Transaction
}

func (c *CoinData) Update(price, histoHour, histoDay bool) {

}

func (c *CoinData) SetName(name string) {
	c.Name = name
}

func (c *CoinData) SetSymbol(symbol string) {
	c.Symbol = symbol
}

func (c *CoinData) SetRank(rank int) {
	c.Rank = rank
}

func (c *CoinData) SetPriceUSD(price float64) {
	c.PriceUSD = price
}

func (c *CoinData) SetTransactions(transactions []*transactions.Transaction) {
	c.transactions = transactions
}

func (c *CoinData) GetTransactions() []*transactions.Transaction {
	return c.transactions
}

func (c *CoinData) AddTransaction(transaction *transactions.Transaction) {
	if c.transactions == nil {
		c.transactions = make([]*transactions.Transaction, 0)
	}

	c.transactions = append(c.transactions, transaction)
}
