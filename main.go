package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/stesla/gotelnet"
)

type config struct {
	Game string
	Host string
	Port int
}

var (
	conn gotelnet.Conn
)

func main() {
	g := flag.String("g", "Discworld", "Set name of game to connect to.")
	h := flag.String("h", "discworld.atuin.net", "Set host to connect to.")
	p := flag.Int("p", 4242, "Set port to connect to.")
	flag.Parse()

	cfg := &config{*g, *h, *p}

	var err error
	conn, err = gotelnet.Dial(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	gui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatal(err)
	}
	defer gui.Close()

	gui.SetManagerFunc(uiLayout)

	if err := uiKeybindings(gui); err != nil {
		log.Fatal(err)
	}

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatal(err)
	}
}
