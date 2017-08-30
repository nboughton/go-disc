package main

import (
	"bufio"
	"fmt"
	"io"

	"github.com/jroimartin/gocui"
)

func recv(g *gocui.Gui, c io.Reader) error {
	bufInput := bufio.NewReader(c)

	for {
		str, _ := bufInput.ReadString('\n')
		g.Update(func(g *gocui.Gui) error {
			v, err := g.View(vMain)
			if err != nil {
				return err
			}

			fmt.Fprint(v, str)
			return nil
		})
	}
}
