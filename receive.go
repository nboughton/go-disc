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
		// Account for if we are in a password prompt or not.
		vI, _ := g.View(vInput)
		if client.PasswordPrompt() {
			vI.Mask = '*'
		} else {
			vI.Mask = 0
		}

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

		vC, err := g.View(vChat)
		if err != nil {
			return err
		}

		// Print to appropriate view
		switch {
		case client.Site.IsChat(line):
			fmt.Fprintln(vC, strings.Replace(line, "  ", "", -1))
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
