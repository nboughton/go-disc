package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/stesla/gotelnet"
)

type config struct {
	Session string
	Host    string
	Port    int
}

var (
	conn gotelnet.Conn
	cfg  *config
)

func main() {
	s := flag.String("s", "Discworld", "Set name of the session.")
	h := flag.String("h", "discworld.atuin.net", "Set host to connect to.")
	p := flag.Int("p", 4242, "Set port to connect to.")
	flag.Parse()

	cfg = &config{*s, *h, *p}

	// Initialise connection
	var err error
	conn, err = gotelnet.Dial(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatal("INIT TELNET ERR:", err)
	}
	defer conn.Close()

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
	go listen(g, conn)

	// Run loop
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatal("GUI ERR:", err)
	}
}
