package transactions

// TransactionType is the type of transaction: buy or sell
type TransactionType int

// TransactionType enumeration
const (
	TransactionBuy TransactionType = iota
	TransactionSell
)

// Transaction describes a but or sell transaction
type Transaction struct {
	ID         int32           `json:"id"`
	Type       TransactionType `json:"type"`
	CoinID     string          `json:"coinId"`
	CoinAmount float64         `json:"coinAmount"`
	DateTime   int64           `json:"dateTime"`
	PriceUSD   float64         `json:"priceUsd"`
}
