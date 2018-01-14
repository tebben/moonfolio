package cmc

// Ticker describes the price of a crypto currency
type Ticker struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUSD         string `json:"price_usd"`
	PriceBTC         string `json:"price_btc"`
	DayVolumeUSD     string `json:"24h_volume_usd"`
	MarketCapUSD     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	MaxSupply        string `json:"max_supply"`
	PercentChange1h  string `json:"percent_change_1h"`
	PercentChange24h string `json:"percent_change_24h"`
	PercentChange7d  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
}

// GlobalData contains the global data retrieved from CMC
type GlobalData struct {
	TotalMarketCapUsd            float64 `json:"total_market_cap_usd"`
	Total24hVolumeUsd            float64 `json:"total_24h_volume_usd"`
	BitcoinPercentageOfMarketCap float64 `json:"bitcoin_percentage_of_market_cap"`
	ActiveCurrencies             int64   `json:"active_currencies"`
	ActiveAssets                 int64   `json:"active_assets"`
	ActiveMarkets                int64   `json:"active_markets"`
	LastUpdated                  int64   `json:"last_updated"`
}

// TickerUpdate contains last retrieved tickers and NextUpdate time
// this struct is used to send over the channel when GetTickersWithUpdates is started
type TickerUpdate struct {
	Tickers    []Ticker
	Error      error
	NextUpdate int64
}
