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

func main() {
	g := flag.String("g", "Discworld", "Set name of game to connect to.")
	h := flag.String("h", "discworld.atuin.net", "Set host to connect to.")
	p := flag.Int("p", 4242, "Set port to connect to.")
	flag.Parse()

	cfg := &config{*g, *h, *p}

	c, err := gotelnet.Dial(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	/*
		go sendHandler(c)


	*/

	gui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatal(err)
	}
	defer gui.Close()

	gui.SetManagerFunc(uiLayout)

	mView, err := gui.View(vMain)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(mView, "test")

	go recvHandler(mView, recv)
	go func() {
		bufInput := bufio.NewReader(c)
		for {
			str, _ := bufInput.ReadString('\n')
			recv <- str
		}
	}()

	if err := gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, uiQuit); err != nil {
		log.Fatal(err)
	}

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatal(err)
	}

}
