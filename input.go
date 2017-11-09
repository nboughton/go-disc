package main

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

func input(g *gocui.Gui, v *gocui.View) error {
	// Trim buffer
	line := strings.TrimSpace(v.Buffer())

	// Write to server connection
	if err := client.Send(line); err != nil {
		return err
	}

	// Insert blank line into the main view
	vM, err := g.View(vMain)
	if err != nil {
		return err
	}
	fmt.Fprintf(vM, "\n")

	// Clear internal buffer and set cursor
	uiZeroLine(v)

	return postSend(line)
}

func postSend(line string) error {
	// Current handling of quit, should probably have conn.Close
	// handled by receiving a CTCP disconnect in the recv func
	if line == "quit" {
		client.Conn.Close()
		return gocui.ErrQuit
	}

	// Add text to tab complete
	if client.LoggedIn() && line != "" {
		for _, s := range strings.Fields(line) {
			dict.Add(s)
		}
	}

	return nil
}
