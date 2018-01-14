package cryptocompare

import (
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
)

const (
	CoinName = "BTC"
)

func TestGetCoinList(t *testing.T) {
	coinList, err := GetCoinList()

	assert.Nil(t, err)
	assert.NotNil(t, coinList)
	assert.NotEmpty(t, coinList.Data)

	btc := coinList.Data[CoinName]
	assert.Equal(t, strings.ToLower(CoinName), strings.ToLower(btc.Name))
}

func TestGetPrice(t *testing.T) {
	price, err := GetPrice(CoinName, []string{"USD", "EUR"}, "", "", false, false)

	assert.Nil(t, err)
	assert.NotNil(t, price)

	if _, ok := price["USD"]; !ok {
		assert.Fail(t, "Expected USD value returned for BTC")
	}

	if _, ok := price["EUR"]; !ok {
		assert.Fail(t, "Expected USD value returned for EUR")
	}
}

func TestGetHistoMinute(t *testing.T) {
	histo, err := GetHistoMinute(CoinName, "USD", "", "", false, false, 2, 10)

	assert.Nil(t, err)
	assert.NotNil(t, histo)
	assert.Len(t, histo.Data, 11, "GetHistoMinute should have returned 10 HistoData objects")

	// test if time difference is 2 minutes
	diff := int(histo.Data[1].Time - histo.Data[0].Time)
	assert.Equal(t, 120, diff)
}

func TestGetHistoHour(t *testing.T) {
	histo, err := GetHistoHour(CoinName, "USD", "", "", false, false, 5, 5)

	assert.Nil(t, err)
	assert.NotNil(t, histo)
	assert.Len(t, histo.Data, 6, "GetHistoHour should have returned 6 HistoData objects")

	// test if time difference is 5 hours
	diff := int(histo.Data[1].Time - histo.Data[0].Time)
	assert.Equal(t, 18000, diff)
}

func TestGetHistoDay(t *testing.T) {
	histo, err := GetHistoDay(CoinName, "USD", "", "", false, false, 2, 4)

	assert.Nil(t, err)
	assert.NotNil(t, histo)
	assert.Len(t, histo.Data, 5, "GetHistoDay should have returned 5 HistoData objects")

	// test if time difference is 2 days
	diff := int(histo.Data[1].Time - histo.Data[0].Time)
	assert.Equal(t, 172800, diff)
}
