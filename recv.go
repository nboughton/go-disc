package main

import (
	"bufio"
	"fmt"
	"io"
	"sync"

	"github.com/jroimartin/gocui"
)

var (
	mu sync.Mutex
)

func recv(g *gocui.Gui, c io.Reader) {
	b := bufio.NewReader(c)

	for {
		mu.Lock()
		str, _, _ := b.ReadLine()
		mu.Unlock()

		g.Update(func(g *gocui.Gui) error {
			v, err := g.View(vMain)
			if err != nil {
				return err
			}

			fmt.Fprintf(v, "%s\n", str)

			return nil
		})

	}
}
