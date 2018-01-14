package ui

import (
	"fmt"
	"time"

	"github.com/jroimartin/gocui"
)

var (
	viewMainBottomName = "mainBottom"
)

func drawMainBottom(g *gocui.Gui) error {
	v, err := g.View(viewMainBottomName)
	if err != nil {
		return err
	}

	v.Clear()

	if timeToNextUpdate <= 0 {
		fmt.Fprintf(v, "updating")
	} else {
		seconds := timeToNextUpdate / 1000
		dt := time.Unix(seconds, 0)
		minute := fmt.Sprintf("%v", dt.Minute())
		if len(minute) < 2 {
			minute = fmt.Sprintf("0%s", minute)
		}

		timeString := fmt.Sprintf("%v:%v", toTimePartString(dt.Minute()), toTimePartString(dt.Second()))
		fmt.Fprintf(v, " %supdate %v", ColorWhite, timeString)
	}

	if cmcError != nil {
		fmt.Fprintf(v, " %s%swoops - %v%s", ColorRed, BoldStart, cmcError, BoldEnd)
	}

	return nil
}

func statusDown() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		return nil
	}
}

func toTimePartString(timePart int) string {
	t := fmt.Sprintf("%v", timePart)
	if len(t) < 2 {
		t = fmt.Sprintf("0%s", t)
	}

	return t
}

func createMainBottom(g *gocui.Gui, x0, y0, x1, y1 int) error {
	if _, err := g.SetView(viewMainBottomName, x0, y0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	return nil
}
