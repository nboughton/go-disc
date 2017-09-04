package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/jroimartin/gocui"
	re "github.com/nboughton/go-utils/regex/common"
)

func listen(g *gocui.Gui, c io.Reader) {
	// Create bufio Reader for incoming data and mutex
	var (
		b  = bufio.NewReader(c)
		mu sync.Mutex
	)

	// Loop input
	for {
		mu.Lock()
		// Bizarrely the other Read* methods don't work
		// ReadLine('\n'), ReadByte('\n') all don't produce
		// output that prints properly into the view window
		var (
			l, _, _ = b.ReadLine()
			line    = handleRecvLine(l)
		)

		// Print new data to view(s)
		g.Update(func(g *gocui.Gui) error {
			v, err := g.View(vMain)
			if err != nil {
				return err
			}

			fmt.Fprintln(v, line)
			mu.Unlock()
			return nil
		})
	}
}

func handleRecvLine(line []byte) string {
	// Trim unwanted characters
	l := strings.TrimPrefix(string(line), "> ")

	// Dont allow logging to the cmdBuffer until
	// after a user is logged in.
	if strings.HasPrefix(l, "You last logged in from") || strings.HasPrefix(l, "You are already playing") {
		cmds.log = true
	}

	// Add words to tab complete dict
	lineNoANSI := re.ANSI.ReplaceAllLiteralString(l, "")
	for _, v := range strings.Fields(lineNoANSI) {
		dict.Add(v)
	}

	return l
}
