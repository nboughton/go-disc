package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/jroimartin/gocui"
	re "github.com/nboughton/go-utils/regex/common"
)

func listen(g *gocui.Gui, c io.Reader) {
	// Create bufio Reader for incoming data and mutex
	var (
		mu sync.Mutex
		b  = bufio.NewReader(c)
	)

	f, _ := os.Create("go-disc.raw.log")
	defer f.Close()

	// Loop input
	for {
		mu.Lock()
		l, _, _ := b.ReadLine()
		fmt.Fprintf(f, "%s\n", l)
		line := handleRecvLine(l)

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
