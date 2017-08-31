package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/jroimartin/gocui"
)

var (
	mu sync.Mutex
)

func recv(g *gocui.Gui, c io.Reader) {
	b := bufio.NewReader(c)

	f, err := os.Create(os.Getenv("HOME") + "/raw_out.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for {
		mu.Lock()
		str, _ := b.ReadString('\n')
		mu.Unlock()

		g.Update(func(g *gocui.Gui) error {
			v, err := g.View(vMain)
			if err != nil {
				return err
			}

			fmt.Fprint(v, "["+str+"]")
			fmt.Fprint(f, str)

			return nil
		})

	}
}
