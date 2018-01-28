package coindata

import (
	"math"
	"sync"
	"time"

	"github.com/tebben/moonfolio/cryptocompare"
	"github.com/tebben/moonfolio/transactions"
)

// Histo contains historical price data
type Histo struct {
	Time     int64
	PriceUSD float64
}

// CoinData contains information
type CoinData struct {
	Name         string
	Symbol       string
	Rank         int
	PriceUSD     float64
	HistoMinute  []*Histo
	HistoDay     []*Histo
	transactions []*transactions.Transaction
	lock         sync.Mutex
}

// SetName to set the name of the coin
func (c *CoinData) SetName(name string) {
	c.Name = name
}

// SetSymbol to set the symbol for the coin
func (c *CoinData) SetSymbol(symbol string) {
	c.Symbol = symbol
}

// SetRank to set the current rank of the coin
func (c *CoinData) SetRank(rank int) {
	c.Rank = rank
}

// SetPriceUSD set the price in USD
func (c *CoinData) SetPriceUSD(price float64) {
	c.PriceUSD = price
}

// SetTransactions sets a slice of transactions for a coin
func (c *CoinData) SetTransactions(transactions []*transactions.Transaction) {
	c.lock.Lock()
	c.transactions = transactions
	c.lock.Unlock()
}

// AddTransaction adds an user transaction to the coin
func (c *CoinData) AddTransaction(transaction *transactions.Transaction) {
	c.lock.Lock()

	if c.transactions == nil {
		c.transactions = make([]*transactions.Transaction, 0)
	}

	c.transactions = append(c.transactions, transaction)
	c.lock.Unlock()
}

// SetHistoDay is used to set historical day data to the coin based on cc histo
func (c *CoinData) SetHistoDay(histoData []*cryptocompare.HistoData) {
	if c.HistoDay == nil {
		c.HistoDay = make([]*Histo, 0)
	}

	for _, h := range histoData {
		nh := &Histo{
			Time:     h.Time * 1000,
			PriceUSD: (h.High + h.Low) / 2,
		}

		c.HistoDay = append([]*Histo{nh}, c.HistoDay...)
	}
}

// SetHistoMinute is used to set historical minute data to the coin based on cc histo
func (c *CoinData) SetHistoMinute(histoData []*cryptocompare.HistoData) {
	for _, h := range histoData {
		nh := &Histo{
			Time:     h.Time * 1000,
			PriceUSD: (h.High + h.Low) / 2,
		}

		c.AddHistoMinute(nh)
	}
}

// AddHistoMinute is used to add historical minute data to the coin, max of 90 minutes back
func (c *CoinData) AddHistoMinute(histo *Histo) {
	if c.HistoMinute == nil {
		c.HistoMinute = make([]*Histo, 0)
	}

	// prepend, curently not sorted, asumming a newer date is added so prepend to slice
	c.HistoMinute = append([]*Histo{histo}, c.HistoMinute...)

	// max to add is 90 minutes = 5400000 ms
	now := int64(time.Now().UnixNano() / int64(time.Millisecond))
	maxEntry := now - 5400000

	for i, h := range c.HistoMinute {
		if h.Time < maxEntry {
			c.HistoMinute = append(c.HistoMinute[:i])
			break
		}
	}
}

// GetTransactions returns all user transactions for a coin
func (c *CoinData) GetTransactions() []*transactions.Transaction {
	if c.transactions == nil {
		return make([]*transactions.Transaction, 0)
	}

	return c.transactions
}

// GetCoinAmount returns the total amount of coins the user is hodling
func (c *CoinData) GetCoinAmount() float64 {
	c.lock.Lock()

	amount := 0.0
	for _, t := range c.GetTransactions() {
		if t.Type == transactions.TransactionBuy {
			amount += t.CoinAmount
		} else if t.Type == transactions.TransactionSell {
			amount -= t.CoinAmount
		}
	}

	c.lock.Unlock()
	return amount
}

// GetBalance returns the Balance in USD for this coin
func (c *CoinData) GetBalance() float64 {
	coinAmount := c.GetCoinAmount()
	return coinAmount * c.PriceUSD
}

func (c *CoinData) GetChange1H() float64 {
	change1H := 0.0

	c.lock.Lock()
	if c.HistoMinute != nil && len(c.HistoMinute) > 0 {
		// find closes but not more than 5 minutes diff
		selectedIndex := -1
		now := int64(time.Now().UnixNano() / int64(time.Millisecond))
		hourBack := now - 3600000

		smallestDiff := float64(-1)

		for i := range c.HistoMinute {
			diff := math.Abs(float64(hourBack) - float64(c.HistoMinute[i].Time))
			if diff < float64(300000) && (smallestDiff == -1 || diff < smallestDiff) {
				selectedIndex = i
				smallestDiff = diff
			}
		}

		if selectedIndex != -1 {
			selectedHisto := c.HistoMinute[selectedIndex]
			change1H = ((c.PriceUSD / selectedHisto.PriceUSD) * 100) - 100
		}
	}
	c.lock.Unlock()

	return change1H
}

func (c *CoinData) GetChange1D() float64 {
	change1D := 0.0

	c.lock.Lock()
	if c.HistoDay != nil && len(c.HistoDay) > 0 {
		firstDay := c.HistoDay[0]
		change1D = ((c.PriceUSD / firstDay.PriceUSD) * 100) - 100
	}
	c.lock.Unlock()

	return change1D
}

func (c *CoinData) GetChange7D() float64 {
	change7D := 0.0

	c.lock.Lock()
	if c.HistoDay != nil && len(c.HistoDay) > 0 {
		lastDay := c.HistoDay[len(c.HistoDay)-1]
		change7D = ((c.PriceUSD / lastDay.PriceUSD) * 100) - 100
	}
	c.lock.Unlock()

	return change7D
}
