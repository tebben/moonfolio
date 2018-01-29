package configuration

// CurrentConfig will be set after loading so it can be accessed from outside
var CurrentConfig Config

// Config contains the application config parameters
type Config struct {
	IntervalSeconds  int                   `json:"intervalSeconds"`
	TransactionsFile string                `json:"transactionsFile"`
	HTTP             TransactionHTTPConfig `json:"http"`
	GUI              GuiConfig             `json:"gui"`
}

// TransactionHTTPConfig holds the information about the transaction HTTP config
type TransactionHTTPConfig struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

// GuiConfig holds the information about the transaction HTTP config
type GuiConfig struct {
	ShowHourPercentage bool `json:"showHourPercentage"`
	ShowDayPercentage  bool `json:"showDayPercentage"`
	ShowWeekPercentage bool `json:"showWeekPercentage"`
	OutlineNumbers     bool `json:"outlineNumbers"`
}
