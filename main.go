package main

import (
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/tebben/moonfolio/cmc"
	"github.com/tebben/moonfolio/configuration"
	"github.com/tebben/moonfolio/ui"
)

var transactions = map[string]float64{
	"bitcoin":     0.0842,
	"vsync-vsx":   11071,
	"hush":        101,
	"ripple":      49.95,
	"cardano":     132.87,
	"iota":        161.91,
	"gulden":      1016.98,
	"verge":       694.30,
	"eccoin":      25000,
	"linda":       10648.25,
	"stronghands": 6917600.41,
}

var (
	updateTicker *time.Ticker
	conf         configuration.Config
	cfgFlag      = flag.String("config", "config.json", "path of the config file")
)

func main() {
	flag.Parse()
	loadConfig()

	createAndStart()
}

func loadConfig() {
	cfg := *cfgFlag

	var err error
	conf, err = configuration.GetConfig(cfg)
	if err != nil {
		log.Fatal("config read error: ", err)
		return
	}

	configuration.SetEnvironmentVariables(&conf)
}

func createAndStart() {
	gui, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Panicln(err)
	}
	defer gui.Close()

	gui.SetManagerFunc(ui.MainLayout)

	if err := gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, ui.Quit); err != nil {
		log.Panicln(err)
	}

	go startCMCTicker()

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func startCMCTicker() {
	time.Sleep(time.Duration(time.Millisecond * 2000))
	cmcTickersChannel, err := cmc.GetTickersWithUpdates(0, cmc.FiatUSD, 0, 0)
	if err != nil {
		return
	}

	go listenForTickers(cmcTickersChannel)
}

func listenForTickers(cmcTickersChannel chan cmc.TickerUpdate) {
	for {
		tickerUpdate := <-cmcTickersChannel

		holdings := ui.Holdings{}
		for _, t := range tickerUpdate.Tickers {
			if amount, ok := transactions[t.ID]; ok {

				coinPrice, _ := strconv.ParseFloat(t.PriceUSD, 64)

				h := ui.HoldingsData{
					CoinName:                 t.Name,
					CoinSymbol:               t.Symbol,
					CoinPrice:                coinPrice,
					HoldingsAmount:           amount,
					HoldingsSymbol:           "$",
					HoldingsBalance:          amount * coinPrice,
					CoinChangePercentageDay:  t.PercentChange24h,
					CoinChangePercentageHour: t.PercentChange1h,
					CoinChangePercentageWeek: t.PercentChange7d,
				}

				holdings = append(holdings, h)
			}
		}

		ui.SetHoldings(holdings, tickerUpdate.Error)

		stopUpdateTimerCountDown()
		go startUpdateTimerCountDown(tickerUpdate.NextUpdate)
	}
}

func startUpdateTimerCountDown(toNextUpdate int64) {
	start := int64(time.Now().UnixNano() / int64(time.Millisecond))
	updateTicker = time.NewTicker(time.Millisecond * time.Duration(200))

	for range updateTicker.C {
		now := int64(time.Now().UnixNano() / int64(time.Millisecond))
		passed := now - start
		left := toNextUpdate - passed
		ui.SetNextUpdateTime(left)
	}
}

func stopUpdateTimerCountDown() {
	if updateTicker != nil {
		updateTicker.Stop()
	}
}
