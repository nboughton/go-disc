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

	// Current handling of quit, should probably have conn.Close
	// handled by receiving a CTCP disconnect in the recv func
	if line == "quit" {
		conn.Close()
		return gocui.ErrQuit
	}

	// Clear internal buffer and set cursor
	v.Clear()
	v.SetOrigin(0, 0)
	v.SetCursor(0, 0)

	// Append line to cmd bufer and set current index to last line
	// Ignore blank returns
	if logToCmdBuffer && line != "" {
		cmdBuffer = append(cmdBuffer, line)
		cmdIdx = len(cmdBuffer)

		// Add it to the autocomplete dict as well
		dict.Add(line)
	}

	return nil
}

func send(c io.Writer, line string) error {
	if _, err := c.Write([]byte(line + "\n")); err != nil {
		return err
	}

	return nil
}
