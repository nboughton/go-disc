package main

import (
	"bufio"
	"fmt"
	"sync"

	"github.com/jroimartin/gocui"
	"github.com/stesla/gotelnet"
)

var mu sync.Mutex

func recv(g *gocui.Gui, c gotelnet.Conn) {
	b := bufio.NewReader(c)
	for {
		mu.Lock()
		str, _ := b.ReadString('\n')
		mu.Unlock()

		g.Update(func(g *gocui.Gui) error {
			v, err := g.View(vMain)
			if err != nil {
				return err
			}

			fmt.Fprint(v, str)
			//fmt.Print(str)
			return nil
		})

	}
}
