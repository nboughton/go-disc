package main

import (
	"bufio"
	"fmt"
	"io"
	//"log"
	//"os"
	"strings"
	"sync"

	"github.com/jroimartin/gocui"
	tc "github.com/nboughton/go-disc/complete"
	re "github.com/nboughton/go-utils/regex/common"
)

var (
	mu             sync.Mutex
	logToCmdBuffer bool
	dict           = tc.New()
)

func listen(g *gocui.Gui, c io.Reader) {
	// Create bufio Reader for incoming data
	b := bufio.NewReader(c)

	// Debugging, lets write raw data to text
	/*
		f, err := os.Create("go-disc.log")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	*/

	// Loop input
	for {
		mu.Lock()
		// Bizarrely the other Read* methods don't work
		// ReadLine('\n'), ReadByte('\n') all don't produce
		// output that prints properly into the view window
		var (
			l, _, _ = b.ReadLine()
			line    = parseRecvLine(l)
		)

		// Print new data to view(s)
		g.Update(func(g *gocui.Gui) error {
			v, err := g.View(vMain)
			if err != nil {
				return err
			}

			fmt.Fprintf(v, "%s\n", line)

			// Print to debug log
			//fmt.Fprintf(f, "%s\n", line)

			mu.Unlock()
			return nil
		})
	}
}

func parseRecvLine(line []byte) string {
	// Trim unwanted characters
	l := strings.TrimPrefix(string(line), "> ")

	// Dont allow logging to the cmdBuffer until
	// after a user is logged in.
	if strings.HasPrefix(l, "You last logged in from") || strings.HasPrefix(l, "You are already playing") {
		logToCmdBuffer = true
	}

	// Add words to tab complete dict
	lineNoANSI := re.ANSI.ReplaceAllLiteralString(l, "")
	f := strings.Fields(lineNoANSI)
	for _, v := range f {
		dict.Add(v)
	}

	return l
}

func copyLine(line []byte) []byte {
	newLine := make([]byte, len(line))
	for i := 0; i < len(line); i++ {
		newLine[i] = line[i]
	}

	return newLine
}
