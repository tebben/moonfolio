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
func GetPrice(fsym ParamFsym, tsyms ParamTsyms, exchange ParamExchange, extraParams ParamExtraparams, sign ParamSign, tryConversion ParamTryConversion) (Price, error) {
	price := priceParse{}
	err := getJSON(buildURI(endpointPrice, fsym, tsyms, exchange, extraParams, sign, tryConversion), &price)
	if err != nil {
		return nil, err
	}

	return price.GetPrices(), nil
}

// GetHistoMinute Get open, high, low, close, volumefrom and volumeto from the each minute historical data.
// This data is only stored for 7 days, if you need more,use the hourly or daily path.
// It uses BTC conversion if data is not available because the coin is not trading in the specified currency
// aggregate	default = 1
// limit		default = 1440, max = 2000
func GetHistoMinute(fsym ParamFsym, tsym ParamTsym, exchange ParamExchange, extraParams ParamExtraparams, sign ParamSign, tryConversion ParamTryConversion, aggregate ParamAggregate, limit ParamLimit) (*Histo, error) {
	return getHisto(endpointHistoMinute, fsym, tsym, exchange, extraParams, sign, tryConversion, aggregate, limit)
}

// GetHistoHour Get open, high, low, close, volumefrom and volumeto from the each hour historical data.
// It uses BTC conversion if data is not available because the coin is not trading in the specified currency.
// aggregate	default = 1
// limit		default = 168, max = 2000
func GetHistoHour(fsym ParamFsym, tsym ParamTsym, exchange ParamExchange, extraParams ParamExtraparams, sign ParamSign, tryConversion ParamTryConversion, aggregate ParamAggregate, limit ParamLimit) (*Histo, error) {
	return getHisto(endpointHistoHour, fsym, tsym, exchange, extraParams, sign, tryConversion, aggregate, limit)
}

// GetHistoDay Get open, high, low, close, volumefrom and volumeto from the each day historical data.
// It uses BTC conversion if data is not available because the coin is not trading in the specified currency.
// aggregate	default = 1
// limit		default = 30, max = 2000
func GetHistoDay(fsym ParamFsym, tsym ParamTsym, exchange ParamExchange, extraParams ParamExtraparams, sign ParamSign, tryConversion ParamTryConversion, aggregate ParamAggregate, limit ParamLimit) (*Histo, error) {
	return getHisto(endpointHistoDay, fsym, tsym, exchange, extraParams, sign, tryConversion, aggregate, limit)
}

func getHisto(endpoint string, fsym ParamFsym, tsym ParamTsym, exchange ParamExchange, extraParams ParamExtraparams, sign ParamSign, tryConversion ParamTryConversion, aggregate ParamAggregate, limit ParamLimit) (*Histo, error) {
	histo := &Histo{}
	err := getJSON(buildURI(endpoint, fsym, tsym, exchange, extraParams, sign, tryConversion, aggregate, limit), histo)
	if err != nil {
		return nil, err
	}

	return histo, nil
}

func getJSON(url string, target interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
