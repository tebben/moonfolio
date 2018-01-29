package runner

import (
	"log"
	"time"

	"github.com/tebben/moonfolio/coindata"

	"github.com/tebben/moonfolio/cryptocompare"
	"github.com/tebben/moonfolio/ui"
)

var (
	updateTicker          *time.Ticker
	countdownTicker       *time.Ticker
	lastHistoDayUpdate    *time.Time
	lastHistoMinuteUpdate *time.Time
	fetchHistoMinuteDone  bool
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
	go updateHisto()
	updateUIHoldings()
}

func updateHisto() {
	if config.GUI.ShowDayPercentage || config.GUI.ShowWeekPercentage {
		updateHistoDay()
	}

	if config.GUI.ShowHourPercentage {
		updateHistoMinute()
	}
}

func updateUIHoldings() {
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
		nowMS := int64(time.Now().UnixNano() / int64(time.Millisecond))

		if p, ok := multi[c.Symbol]; ok {
			priceUSD := p["USD"]
			c.SetPriceUSD(priceUSD)

			if fetchHistoMinuteDone && config.GUI.ShowHourPercentage {
				{
					histo := &coindata.Histo{
						Time:     nowMS,
						PriceUSD: priceUSD,
					}
					c.AddHistoMinute(histo)
				}
			}
		}
	}
}

func updateHistoDay() {
	now := time.Now()
	if lastHistoDayUpdate == nil || (lastHistoDayUpdate.Year() < now.Year() || lastHistoDayUpdate.Month() < now.Month() || lastHistoDayUpdate.Day() < now.Day()) {
		lastHistoDayUpdate = &now

		for k := range userCoins {
			histoDay, err := cryptocompare.GetHistoDay(cryptocompare.ParamFsym(userCoins[k].Symbol), "USD", "", "", false, false, 1, 7)
			if err != nil {
				// Todo: set error
			}

			userCoins[k].SetHistoDay(histoDay.Data)
			go updateUIHoldings()
			time.Sleep(time.Millisecond * 1000)
		}
	}
}

func updateHistoMinute() {
	now := time.Now()
	if lastHistoMinuteUpdate == nil {
		lastHistoMinuteUpdate = &now

		for k := range userCoins {
			histoMinute, err := cryptocompare.GetHistoMinute(cryptocompare.ParamFsym(userCoins[k].Symbol), "USD", "", "", false, false, 1, 60)
			if err != nil {
				// Todo: set error
			}

			userCoins[k].SetHistoMinute(histoMinute.Data)
			go updateUIHoldings()
			time.Sleep(time.Millisecond * 1000)
		}

		fetchHistoMinuteDone = true
	}
}
