package cryptocompare

// CoinList is the response returned by cryptocompare on the /coins endpoint
type CoinList struct {
	Response     string          `json:"Response"`
	Message      string          `json:"Message"`
	BaseImageURL string          `json:"BaseImageUrl"`
	BaseLinkURL  string          `json:"BaseLinkUrl"`
	Type         int             `json:"Type"`
	Data         map[string]Coin `json:"Data"`
}

// Coin describes a coin
type Coin struct {
	ID        string `json:"Id"`
	URL       string `json:"Url"`
	ImageURL  string `json:"ImageUrl"`
	Name      string `json:"Name"`
	Symbol    string `json:"Symbol"`
	CoinName  string `json:"CoinName"`
	FullName  string `json:"FullName"`
	Algorithm string `json:"Algorithm"`
	ProofType string `json:"ProofType"`
	SortOrder string `json:"SortOrder"`
}

// Price contains price information of a coin
type Price map[string]float64

// priceParse used to parse the prices into, ,need because of the %@#^&% generic fields. GAWT
type priceParse map[string]interface{}

func (p *priceParse) GetPrices() Price {
	price := Price{}

	for k, v := range *p {
		price[k] = v.(float64)
	}

	return price
}

// Histo response from endpoints HistoMinute, HistoHour, HistoDay
type Histo struct {
	Response          string         `json:"Response"`
	Type              int            `json:"Type"`
	Aggregated        bool           `json:"Aggregated"`
	TimeTo            int64          `json:"TimeTo"`
	TimeFrom          int64          `json:"TimeFrom"`
	FirstValueInArray bool           `json:"FirstValueInArray"`
	ConversionType    ConversionType `json:"ConversionType"`
	Data              []HistoData    `json:"Data"`
}

// ConversionType something that is returned
type ConversionType struct {
	Type             string `json:"type"`
	ConversionSymbol string `json:"conversionSymbol"`
}

// HistoData contains open, high, low, close, volumefrom and volumeto from historical data
type HistoData struct {
	Time       int64   `json:"time"`
	Close      float64 `json:"close"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	Open       float64 `json:"open"`
	Volumefrom float64 `json:"volumefrom"`
	Volumeto   float64 `json:"volumeto"`
}
