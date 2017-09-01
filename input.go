package main

import (
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

	// Current handling of quit, should really have conn.Close
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
	cmdBuffer = append(cmdBuffer, line)
	cmdIdx = len(cmdBuffer)

	return nil
}

func send(c io.Writer, line string) error {
	if _, err := c.Write([]byte(line + "\n")); err != nil {
		return err
	}

	return nil
}
