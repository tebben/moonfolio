package ui

import (
	"fmt"
	"sort"
	"strings"

	"github.com/tebben/moonfolio/coindata"

	"github.com/jroimartin/gocui"
)

var (
	viewOverviewName = "mainOverview"
)

func drawMainOverview(g *gocui.Gui) error {
	v, err := g.View(viewOverviewName)
	if err != nil {
		return err
	}
	xSize, _ := v.Size()
	//v.BgColor = gocui.ColorBlack
	v.Clear()
	fmt.Fprintln(v, createHeader())
	fmt.Fprintln(v, getColumnHeadSpacer(xSize))

	// Sort holdings curently based on holding balance
	sort.Sort(holdings)

	for i := range holdings {
		h := holdings[i]
		fmt.Fprintln(v, createDataRow(h))
		fmt.Fprintln(v, getColumnSpacer(xSize))
	}

	return nil
}

func createHeader() string {
	textColor := ColorWhite
	data := []columnText{
		columnText{Length: 19, Text: "NAME", Styling: []string{textColor, BoldStart}},
		columnText{Length: 14, Text: "PRICE", Styling: []string{textColor, BoldStart}},
		columnText{Length: 10, Text: fmt.Sprintf("1H%s", "%"), Styling: []string{textColor, BoldStart}},
		columnText{Length: 10, Text: fmt.Sprintf("1D%s", "%"), Styling: []string{textColor, BoldStart}},
		columnText{Length: 10, Text: fmt.Sprintf("7D%s", "%"), Styling: []string{textColor, BoldStart}},
		columnText{Length: 13, Text: "AMOUNT", Styling: []string{textColor, BoldStart}},
		columnText{Length: 13, Text: "BALANCE", Styling: []string{textColor, BoldStart}},
	}

	return createColumnString(data)
}

func createDataRow(h *coindata.CoinData) string {
	data := []columnText{
		columnText{Length: 19, Text: fmt.Sprintf("%s (%s)", h.Name, h.Symbol), Styling: []string{ColorGray, BoldStart}},
		columnText{Length: 14, Text: fmt.Sprintf("%s%v", "$", h.PriceUSD), Styling: []string{ColorGray, BoldStart}},
		columnText{Length: 10, Text: fmt.Sprintf("%s%s", h.GetChange1H(), "%"), Styling: []string{getChangeColorStyle(h.GetChange1H()), BoldStart}},
		columnText{Length: 10, Text: fmt.Sprintf("%s%s", h.GetChange1D(), "%"), Styling: []string{getChangeColorStyle(h.GetChange1D()), BoldStart}},
		columnText{Length: 10, Text: fmt.Sprintf("%s%s", h.GetChange7D(), "%"), Styling: []string{getChangeColorStyle(h.GetChange7D()), BoldStart}},
		columnText{Length: 13, Text: fmt.Sprintf("%v", h.GetCoinAmount()), Styling: []string{ColorGray, BoldStart}},
		columnText{Length: 13, Text: fmt.Sprintf("%s%.2f", "$", h.GetBalance()), Styling: []string{ColorGray, BoldStart}},
	}

	return createColumnString(data)
}

func createMainOverview(g *gocui.Gui, x0, y0, x1, y1 int) error {
	if _, err := g.SetView(viewOverviewName, x0, y0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	return nil
}

func getChangeColorStyle(change string) string {
	s := ColorGreen
	if strings.HasPrefix(change, "-") {
		s = ColorRed
	}

	return s
}
