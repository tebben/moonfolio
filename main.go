package main

import (
	"flag"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/tebben/moonfolio/configuration"
	"github.com/tebben/moonfolio/runner"
	"github.com/tebben/moonfolio/ui"
)

var (
	conf    configuration.Config
	cfgFlag = flag.String("config", "config.json", "path of the config file")
	gui     *gocui.Gui
)

func main() {
	flag.Parse()
	loadConfig()
	createAndStart()
}

func loadConfig() {
	var err error

	cfg := *cfgFlag
	conf, err = configuration.GetConfig(cfg)
	if err != nil {
		log.Fatal("config read error: ", err)
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

	go runner.Start(&conf, gui)

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
