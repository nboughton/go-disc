package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/jroimartin/gocui"
)

func input(g *gocui.Gui, v *gocui.View) error {
	// Trim buffer
	line := strings.TrimSpace(v.Buffer())

	// Write to server connection
	if err := send(conn, line); err != nil {
		return err
	}

	// Insert blank line into the main view
	vM, err := g.View(vMain)
	if err != nil {
		return err
	}
	fmt.Fprintf(vM, "\n")

	// Clear internal buffer and set cursor
	zeroLine(v)

	if err := handlePostSend(line); err != nil {
		return err
	}

	return nil
}

func handlePostSend(line string) error {
	// Current handling of quit, should probably have conn.Close
	// handled by receiving a CTCP disconnect in the recv func
	if line == "quit" {
		conn.Close()
		return gocui.ErrQuit
	}

	// Log sent line
	cmds.Log(line)

	// Add text to tab complete
	if cmds.Logging() && line != "" {
		for _, s := range strings.Fields(line) {
			dict.Add(s)
		}
	}

	return nil
}

func send(c io.Writer, line string) error {
	if _, err := c.Write([]byte(line + "\n")); err != nil {
		return err
	}

	return nil
}
