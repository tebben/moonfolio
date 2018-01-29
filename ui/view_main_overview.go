package ui

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/tebben/moonfolio/coindata"

	"github.com/jroimartin/gocui"
)

const (
	columnName int = iota
	columnPrice
	columnChange1H
	columnChange1D
	columnChange7D
	columnAmount
	columnBalance
)

var (
	viewOverviewName = "mainOverview"
	columns          = map[int]columnDefinition{
		columnName:     columnDefinition{Header: "NAME", Size: 25, StylingHeader: getRowHeaderStyling, StylingBody: getRowColumnDefaultStyling, Enabled: true},
		columnPrice:    columnDefinition{Header: "PRICE", Size: 14, StylingHeader: getRowHeaderStyling, StylingBody: getRowColumnDefaultStyling, Enabled: true},
		columnChange1H: columnDefinition{Header: fmt.Sprintf("1H %s", "%"), Size: 10, StylingHeader: getRowHeaderStyling, StylingBody: getRowChangeStyling, Enabled: true},
		columnChange1D: columnDefinition{Header: fmt.Sprintf("1D %s", "%"), Size: 10, StylingHeader: getRowHeaderStyling, StylingBody: getRowChangeStyling, Enabled: true},
		columnChange7D: columnDefinition{Header: fmt.Sprintf("7D %s", "%"), Size: 10, StylingHeader: getRowHeaderStyling, StylingBody: getRowChangeStyling, Enabled: true},
		columnAmount:   columnDefinition{Header: "AMOUNT", Size: 13, StylingHeader: getRowHeaderStyling, StylingBody: getRowColumnDefaultStyling, Enabled: true},
		columnBalance:  columnDefinition{Header: "BALANCE", Size: 13, StylingHeader: getRowHeaderStyling, StylingBody: getRowColumnDefaultStyling, Enabled: true},
	}
)

func drawMainOverview(g *gocui.Gui) error {
	setEnabledColumns()

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

func setEnabledColumns() {
	hour := columns[columnChange1H]
	hour.Enabled = Config.ShowHourPercentage

	day := columns[columnChange1D]
	day.Enabled = Config.ShowDayPercentage

	week := columns[columnChange7D]
	week.Enabled = Config.ShowWeekPercentage
}

func createHeader() string {
	data := []columnText{}

	var keys []int
	for k := range columns {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		r := columns[k]
		if !r.Enabled {
			continue
		}

		data = append(data, columnText{Length: r.Size, Text: r.Header, Styling: r.StylingHeader()})
	}

	return createColumnString(data)
}

func createDataRow(h *coindata.CoinData) string {
	data := []columnText{}
	data = append(data, columnText{Length: columns[columnName].Size, Text: h.Name, Styling: columns[columnName].StylingBody(nil)})
	data = append(data, columnText{Length: columns[columnPrice].Size, Text: fmt.Sprintf("%s%v", "$", h.PriceUSD), Styling: columns[columnPrice].StylingBody(nil)})

	if columns[columnChange1H].Enabled {
		data = append(data, columnText{Length: columns[columnChange1H].Size, Text: fmt.Sprintf("%s", getChange(h.GetChange1H)), Styling: columns[columnChange1H].StylingBody(h.GetChange1H())})
	}

	if columns[columnChange1D].Enabled {
		data = append(data, columnText{Length: columns[columnChange1D].Size, Text: fmt.Sprintf("%s", getChange(h.GetChange1D)), Styling: columns[columnChange1D].StylingBody(h.GetChange1D())})
	}

	if columns[columnChange7D].Enabled {
		data = append(data, columnText{Length: columns[columnChange7D].Size, Text: fmt.Sprintf("%s", getChange(h.GetChange7D)), Styling: columns[columnChange7D].StylingBody(h.GetChange7D())})
	}

	data = append(data, columnText{Length: columns[columnAmount].Size, Text: fmt.Sprintf("%v", floatToString(h.GetCoinAmount())), Styling: columns[columnAmount].StylingBody(nil)})
	data = append(data, columnText{Length: columns[columnBalance].Size, Text: fmt.Sprintf("%s%.2f", "$", h.GetBalance()), Styling: columns[columnBalance].StylingBody(nil)})

	return createColumnString(data)
}

func getChange(f func() float64) string {
	val := f()
	if val == coindata.Empty {
		return "-"
	}

	return fmt.Sprintf("%.2f%s", val, "%")
}

func floatToString(input float64) string {
	return strconv.FormatFloat(input, 'f', 4, 64)
}

func createMainOverview(g *gocui.Gui, x0, y0, x1, y1 int) error {
	if _, err := g.SetView(viewOverviewName, x0, y0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	return nil
}

func getChangeColorStyle(change float64) string {
	if change == coindata.Empty {
		return ColorWhite
	}

	s := ColorGreen
	if change < 0 {
		s = ColorRed
	}

	return s
}
