package cmc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	endpointAPI    = "https://api.coinmarketcap.com/v1"
	endpointTicker = fmt.Sprintf("%s/ticker/", endpointAPI)
	endpointGlobal = fmt.Sprintf("%s/global/", endpointAPI)
)

var (
	client                = &http.Client{Timeout: 30 * time.Second}
	cmcTickersChannel     = make(chan TickerUpdate)
	getTicker             *time.Ticker
	minUpdateMilliseconds = int64(300000)
)

// GetTickers returns a list of tickers.
// Set start and limit to 0 to get all tickers from CoinMarketCap
func GetTickers(fiat Fiat, start, limit int) ([]Ticker, error) {
	currencies := []Ticker{}
	err := getJSON(fmt.Sprintf("%s?convert=%s&start=%v&limit=%v", endpointTicker, fiat.String(), start, limit), &currencies)
	if err != nil {
		return nil, err
	}

	return currencies, nil
}

// GetTickerForCurrency returns a ticker for specific
// currency supplied by the caller
// ToDo: Check if currency is available
func GetTickerForCurrency(currency string, fiat Fiat) ([]Ticker, error) {
	currencies := []Ticker{}
	err := getJSON(fmt.Sprintf("%s/%s/?convert=%s", endpointTicker, currency, fiat.String()), &currencies)
	if err != nil {
		return nil, err
	}

	return currencies, nil
}

// GetGlobalData retrieves the global data from CMC such as total market cap, bitcoin percentage of market cap etc
func GetGlobalData(fiat Fiat) (*GlobalData, error) {
	globalData := GlobalData{}
	err := getJSON(fmt.Sprintf("%s?convert=%s", endpointGlobal, fiat.String()), &globalData)
	if err != nil {
		return nil, err
	}

	return &globalData, nil
}

// GetTickersWithUpdates gets the tickers from CMC and sends it over the returned channel
//  updateMs: 	interval to retrieve new data, set to 0 to use the fastest update frequency (curently 5 minutes)
// 				if the supplied milliseconds are lower than the set minUpdateMilliseconds, it will be set to minUpdateMilliseconds
//  fiat:		Fiat enum if you want to get some extra fields in a certain fiat currency
//  start:		return results from rank [start] and above, use start = 0, limit = 0 to retrieve all tickers
//  limit:		return a maximum of [limit] results, use start = 0, limit = 0 to retrieve all tickers
func GetTickersWithUpdates(updateMs int64, fiat Fiat, start, limit int) (chan TickerUpdate, error) {
	go func() {
		updateFrequency := updateMs
		if updateFrequency < minUpdateMilliseconds {
			updateFrequency = minUpdateMilliseconds
		}

		go getTickersAndSendOverChannel(fiat, start, limit, updateFrequency)

		getTicker = time.NewTicker(time.Millisecond * time.Duration(updateFrequency))
		for range getTicker.C {
			getTickersAndSendOverChannel(fiat, start, limit, updateFrequency)
		}
	}()

	return cmcTickersChannel, nil
}

// GetTickersWithUpdatesStop stops retrieving tickers every xx seconds
func GetTickersWithUpdatesStop() {
	if getTicker != nil {
		getTicker.Stop()
	}
}

func getTickersAndSendOverChannel(fiat Fiat, start, limit int, nextUpdate int64) {
	tickers, err := GetTickers(fiat, start, limit)
	tickerUpdate := TickerUpdate{
		Tickers:    tickers,
		NextUpdate: nextUpdate,
		Error:      err,
	}

	cmcTickersChannel <- tickerUpdate
}

func getJSON(url string, target interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
