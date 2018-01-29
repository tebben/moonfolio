package ui

import (
	"github.com/jroimartin/gocui"
	"github.com/tebben/moonfolio/configuration"
)

var (
	holdings         = make(Holdings, 0)
	timeToNextUpdate int64
	cmcError         error
	gui              *gocui.Gui
	Config           configuration.GuiConfig
)

// ReDraw the entire layout
func ReDraw() {
	gui.Update(func(g *gocui.Gui) error {
		drawMainTop(g)
		drawMainOverview(g)
		drawMainBottom(g)

		return nil
	})
}

// SetHoldings sets the current holding data based on transactions and cryptocompare tickers
func SetHoldings(h Holdings, err error) {
	if h != nil {
		holdings = h
	}

	cmcError = err
	ReDraw()
}

// SetNextUpdateTime sets the time left to next coin update is fetched
func SetNextUpdateTime(nextUpdate int64) {
	timeToNextUpdate = nextUpdate
	redrawBottom()
}

func redrawBottom() {
	gui.Update(func(g *gocui.Gui) error {
		drawMainBottom(g)

		return nil
	})
}

// MainLayout creates the layout for the gui
func MainLayout(g *gocui.Gui) error {
	g.BgColor = gocui.ColorBlack
	gui = g
	maxX, maxY := g.Size()
	bottomHeight := 2
	topHeight := len(GetLogoV2()) + 1

	//err := createTop(g, -1, -1, maxX, int(0.17*float32(maxY)))
	err := createMainTop(g, -1, -1, maxX, topHeight)
	if err != nil {
		return err
	}

	err = createMainOverview(g, -1, topHeight, maxX, maxY-bottomHeight)
	if err != nil {
		return err
	}

	err = createMainBottom(g, -1, maxY-bottomHeight, maxX, maxY)
	if err != nil {
		return err
	}

	return nil
}

// Quit stops the gui
func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
