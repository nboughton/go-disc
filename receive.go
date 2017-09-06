package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/jroimartin/gocui"
	re "github.com/nboughton/go-utils/regex/common"
)

var mu sync.Mutex

func listen(g *gocui.Gui) {
	for line := range client.Receive() {
		mu.Lock()
		printToViews(g, processLine(line))
	}
}

func printToViews(g *gocui.Gui, line string) {
	// Print new data to view(s)
	g.Update(func(g *gocui.Gui) error {
		// Get views to print to
		vM, err := g.View(vMain)
		if err != nil {
			return err
		}

		vT, err := g.View(vTop)
		if err != nil {
			return err
		}

		// Print to appropriate view
		switch {
		case client.Site.IsChat(line):
			fmt.Fprintln(vT, line)
		default:
			fmt.Fprintln(vM, line)
		}

		mu.Unlock()
		return nil
	})
}

func processLine(line string) string {
	// Trim unwanted characters
	l := strings.TrimPrefix(line, "> ")

	// Add words to tab complete dict
	lineNoANSI := re.ANSI.ReplaceAllLiteralString(l, "")
	for _, v := range strings.Fields(lineNoANSI) {
		dict.Add(v)
	}

	return l
}
