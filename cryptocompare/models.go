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
