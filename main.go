package main

import (
	"flag"
	//"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/nboughton/go-disc/mud"
	//"github.com/stesla/gotelnet"
)

func main() {
	s := flag.String("s", "Discworld", "Set name of the session.")
	h := flag.String("h", "discworld.atuin.net", "Set host to connect to.")
	p := flag.Int("p", 4242, "Set port to connect to.")
	flag.Parse()

	cfg = &config{*s, *h, *p}

	// Initialise connection
	var err error
	client, err = mud.NewClient(*h, *p)
	if err != nil {
		log.Fatal(err)
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
