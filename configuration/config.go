package configuration

// CurrentConfig will be set after loading so it can be accessed from outside
var CurrentConfig Config

// Config contains the application config parameters
type Config struct {
	IntervalSeconds  int                   `json:"intervalSeconds"`
	TransactionsPath string                `json:"transactionsPath"`
	HTTP             TransactionHTTPConfig `json:"http"`
}

// TransactionHTTPConfig holds the information about the transaction HTTP config
type TransactionHTTPConfig struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}
