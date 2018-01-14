package cryptocompare

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	endpointAPI                  = "https://min-api.cryptocompare.com/data"
	endpointCoinList             = "https://www.cryptocompare.com/api/data/coinlist/"
	endpointPrice                = fmt.Sprintf("%s/price", endpointAPI)
	endpointPriceMulti           = fmt.Sprintf("%s/pricemulti", endpointAPI)
	endpointPriceHistorical      = fmt.Sprintf("%s/pricehistorical", endpointAPI)
	endpointCoinSnapshot         = fmt.Sprintf("%s/coinsnapshot/", endpointAPI)
	endpointCoinSnapshotFullByID = fmt.Sprintf("%s/coinsnapshotfullbyid/", endpointAPI)
	endpointSocialStats          = fmt.Sprintf("%s/socialstats/", endpointAPI)
	endpointHistoMinute          = fmt.Sprintf("%s/histominute", endpointAPI)
	endpointHistoHour            = fmt.Sprintf("%s/histohour", endpointAPI)
	endpointHistoDay             = fmt.Sprintf("%s/histoday", endpointAPI)
)

var (
	client = &http.Client{Timeout: 30 * time.Second}
)

// GetCoinList returns general info for all the coins available on cryptocompare.
func GetCoinList() (*CoinList, error) {
	coinList := &CoinList{}
	err := getJSON(endpointCoinList, coinList)
	if err != nil {
		return nil, err
	}

	return coinList, nil
}

// GetPrice returns the latest price for a list of one or more currencies.
func GetPrice(fsym string, tsyms []string, optionals *GetPriceOptionals) (Price, error) {
	price := priceParse{}
	err := getJSON(buildURI(endpointPrice, optionals, &Fsym{fsym}, &Tsyms{tsyms}), &price)
	if err != nil {
		return nil, err
	}

	return price.GetPrices(), nil
}

func getJSON(url string, target interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
