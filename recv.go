package main

import (
	"bufio"
	"fmt"
	"io"
	"sync"

	"github.com/jroimartin/gocui"
)

var mu sync.Mutex

func recv(gui *gocui.Gui, c io.Reader) {
	b := bufio.NewReader(c)
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
			//v.Write([]byte(str))
			//fmt.Print(str)
			return nil
		})
	}
}
