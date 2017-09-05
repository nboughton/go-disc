package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//"regexp"
	"strings"
	"sync"

	"github.com/jroimartin/gocui"
	re "github.com/nboughton/go-utils/regex/common"
)

var (
	ansiEOL      = "[39;49m[0;10m"
	dwAnsiTalker = "[1m[32m"
	mu           sync.Mutex
)

func listen(g *gocui.Gui, c io.Reader) {
	// Create bufio Reader for incoming data
	b := bufio.NewReader(c)

	// Open file and print raw data for testing
	f, _ := os.Create("go-disc.raw.log")
	defer f.Close()

	// Loop input
	for {
		mu.Lock()
		l, _, _ := b.ReadLine()
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

func processLine(line []byte) string {
	// Trim unwanted characters
	l := strings.TrimPrefix(string(line), "> ")

	// Dont allow logging to the cmd history until
	// after a user is logged in.
	if strings.HasPrefix(l, "You last logged in from") || strings.HasPrefix(l, "You are already playing") {
		cmds.SetLogging(true)
	}

	// Add words to tab complete dict
	lineNoANSI := re.ANSI.ReplaceAllLiteralString(l, "")
	for _, v := range strings.Fields(lineNoANSI) {
		dict.Add(v)
	}

	return l
}
