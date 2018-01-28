package runner

import (
	"log"
	"time"

	"github.com/tebben/moonfolio/coindata"

	"github.com/tebben/moonfolio/cryptocompare"
	"github.com/tebben/moonfolio/ui"
)

var (
	updateTicker    *time.Ticker
	countdownTicker *time.Ticker
)

func startUpdater(updateMs int64) {
	// start with an update
	runUpdate(updateMs)

	updateTicker = time.NewTicker(time.Millisecond * time.Duration(updateMs))
	for range updateTicker.C {
		runUpdate(updateMs)
	}
}

func runUpdate(updateMs int64) {
	go update()
	stopCountdownTicker()
	go startCountdownTicker(updateMs)
}

func stopUpdater() {
	if updateTicker != nil {
		updateTicker.Stop()
	}
}

func startCountdownTicker(toNextUpdate int64) {
	start := int64(time.Now().UnixNano() / int64(time.Millisecond))
	countdownTicker = time.NewTicker(time.Millisecond * time.Duration(200))

	for range countdownTicker.C {
		now := int64(time.Now().UnixNano() / int64(time.Millisecond))
		passed := now - start
		left := toNextUpdate - passed
		ui.SetNextUpdateTime(left)
	}
}

func stopCountdownTicker() {
	if countdownTicker != nil {
		countdownTicker.Stop()
	}
}

func update() {
	updatePrice()
	updateHisto()

	// set data to ui
	holdings := make(ui.Holdings, 0)
	for k := range userCoins {
		copy := &coindata.CoinData{}
		*copy = *userCoins[k]
		holdings = append(holdings, copy)
	}

	ui.SetHoldings(holdings, nil)
}

func updatePrice() {
	fsyms := make([]string, 0)
	for _, c := range userCoins {
		fsyms = append(fsyms, c.Symbol)
	}

	multi, err := cryptocompare.GetPriceMulti(fsyms, []string{"USD"}, "", "", false, false)
	if err != nil {
		//ToDo: set error, remove fatal
		log.Fatalf("Error getting price multi: %v", err)
	}

	// set the price for a user coin from retrieved price data
	for _, c := range userCoins {
		if p, ok := multi[c.Symbol]; ok {
			c.SetPriceUSD(p["USD"])
		}
	}
}

func updateHisto() {

}

// Every 00:00 gmt get price histo day to set 1 day and 7 d
// every
