package main

import (
	"strings"

	"github.com/jroimartin/gocui"
)

func send(g *gocui.Gui, v *gocui.View) error {
	line := strings.TrimSpace(v.Buffer())

	if _, err := conn.Write([]byte(line + "\n")); err != nil {
		return err
	}

	return nil
}
