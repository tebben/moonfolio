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
	ID         int             `json:"id"`
	Type       TransactionType `json:"type"`
	CoinID     string          `json:"coinId"`
	CoinAmount float64         `json:"coinAmount"`
	DateTime   int64           `json:"dateTime"`
	PriceUSD   float64         `json:"priceUsd"`
	Hide       bool            `json:"hide"`
}

// IStore defines the operations a store must implement
type IStore interface {
	GetTransactions() ([]*Transaction, error)
	AddTransaction(t *Transaction) error
	UpdateTransaction(t *Transaction) error
	DeleteTransaction(t *Transaction) error
}
