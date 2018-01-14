package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

var (
	viewTopName = "mainTop"
)

func drawMainTop(g *gocui.Gui) error {
	v, err := g.View(viewTopName)
	if err != nil {
		return err
	}

	v.Clear()

	logo := GetLogoV2()
	for _, s := range logo {
		fmt.Fprintln(v, s)
	}

	total := 0.0
	for _, h := range holdings {
		total = total + h.HoldingsBalance
	}

	fmt.Fprintf(v, fmt.Sprintf("total: $%.2f", total))

	return nil
}

func createMainTop(g *gocui.Gui, x0, y0, x1, y1 int) error {
	if _, err := g.SetView(viewTopName, x0, y0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	return nil
}
