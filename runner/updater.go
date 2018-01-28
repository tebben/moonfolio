package runner

import (
	"time"

	"github.com/tebben/moonfolio/ui"
)

var (
	updateTicker    *time.Ticker
	countdownTicker *time.Ticker
)

func startUpdater(updateMs int64) {
	updateTicker = time.NewTicker(time.Millisecond * time.Duration(updateMs))

	for range updateTicker.C {
		update()
		stopUpdateTimerCountDown()
		startUpdateTimerCountDown(updateMs)
	}
}

func stopUpdater() {
	if updateTicker != nil {
		updateTicker.Stop()
	}
}

func startUpdateTimerCountDown(toNextUpdate int64) {
	start := int64(time.Now().UnixNano() / int64(time.Millisecond))
	countdownTicker = time.NewTicker(time.Millisecond * time.Duration(200))

	for range countdownTicker.C {
		now := int64(time.Now().UnixNano() / int64(time.Millisecond))
		passed := now - start
		left := toNextUpdate - passed
		ui.SetNextUpdateTime(left)
	}
}

func stopUpdateTimerCountDown() {
	if countdownTicker != nil {
		countdownTicker.Stop()
	}
}
