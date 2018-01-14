package ui

type Holdings []HoldingsData

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
	return d[i].HoldingsBalance > d[j].HoldingsBalance
}

type HoldingsData struct {
	CoinName                 string
	CoinSymbol               string
	CoinRank                 int
	CoinPrice                float64
	CoinChangePercentageHour string
	CoinChangePercentageDay  string
	CoinChangePercentageWeek string

	HoldingsSymbol  string
	HoldingsAmount  float64
	HoldingsBalance float64
	HoldingsCost    float64
}
