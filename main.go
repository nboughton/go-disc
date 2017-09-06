package main

import (
	"flag"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/nboughton/go-disc/complete"
	"github.com/nboughton/go-disc/mud"
)

var (
	client *mud.Client
	dict   = complete.New()
)

func main() {
	s := flag.String("s", "Discworld", "Set name of the session.")
	flag.Parse()

	// Initialise connection
	var err error
	client, err = mud.NewClient(*s)
	if err != nil {
		log.Fatal("CONNECTION ERR:", err)
	}

	// Initialise g
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		log.Fatal("INIT GUI ERR:", err)
	}
	defer g.Close()

	g.Cursor = true

	// Set layout manager
	g.SetManagerFunc(uiLayout)

	// Set keybindings
	if err := uiKeybindings(g); err != nil {
		log.Fatal("KBDG ERR:", err)
	}

	// Set up receiver from mud server
	go listen(g)

	// Run loop
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatal("GUI ERR:", err)
	}
}
