package main

import (
	"bufio"
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

	// Initialise connection
	var err error
	conn, err = gotelnet.Dial(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatal("INIT TELNET ERR:", err)
	}
	defer conn.Close()

	// Initialise gui
	gui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatal("INIT GUI ERR:", err)
	}
	defer gui.Close()

	// Set layout manager
	gui.SetManagerFunc(uiLayout)

	// Set keybindings
	if err := uiKeybindings(gui); err != nil {
		log.Fatal("KBDG ERR:", err)
	}

	// Set up receiver from mud server
	//go recv(gui, conn)
	go func() {
		b := bufio.NewReader(conn)
		for {
			mu.Lock()
			str, _ := b.ReadString('\n')
			mu.Unlock()

			gui.Update(func(g *gocui.Gui) error {
				v, err := g.View(vMain)
				if err != nil {
					return err
				}

				fmt.Fprint(v, str)
				return nil
			})
		}
	}()

	// Run loop
	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatal("GUI ERR:", err)
	}
}
