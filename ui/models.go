package ui

import (
	"github.com/tebben/moonfolio/coindata"
)

type Holdings []*coindata.CoinData

// Len is part of sort.Interface.
func (d Holdings) Len() int {
	return len(d)
}

// Swap is part of sort.Interface.
func (d Holdings) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Less is part of sort.Interface. We use count as the value to sort by
func (d Holdings) Less(i, j int) bool {
	return d[i].GetBalance() > d[j].GetBalance()
}
