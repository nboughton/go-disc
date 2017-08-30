package main

import (
	"bufio"
	"fmt"
	"io"

	"github.com/jroimartin/gocui"
)

func recv(g *gocui.Gui, c io.Reader) error {
	bufInput := bufio.NewScanner(c)

	for bufInput.Scan() {
		g.Update(func(g *gocui.Gui) error {
			v, err := g.View(vMain)
			if err != nil {
				return err
			}

			fmt.Fprint(v, bufInput.Text())
			return nil
		})
	}

	return nil
}
