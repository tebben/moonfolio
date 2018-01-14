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
	price, err := GetPrice(CoinName, []string{"USD", "EUR"}, nil)

	assert.Nil(t, err)
	assert.NotNil(t, price)

	if _, ok := price["USD"]; !ok {
		assert.Fail(t, "Expected USD value returned for BTC")
	}

	if _, ok := price["EUR"]; !ok {
		assert.Fail(t, "Expected USD value returned for EUR")
	}
}
