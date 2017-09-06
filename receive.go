package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/jroimartin/gocui"
	re "github.com/nboughton/go-utils/regex/common"
)

var (
	//ansiEOL      = "[39;49m[0;10m"
	mu           sync.Mutex
	dwAnsiTalker = "[1m[32m"
)

func listen(g *gocui.Gui) {
	// Open file and print raw data for testing
	f, _ := os.Create("go-disc.raw.log")
	defer f.Close()

	// Loop input
	for l := range client.Receive() {
		mu.Lock()
		line := processLine(l)

		// Print debugging data to raw log
		fmt.Fprintln(f, line)

		printToViews(g, line)
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
		case strings.Contains(line, dwAnsiTalker):
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
