package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"

	"github.com/nboughton/go-utils/input"
	"github.com/stesla/gotelnet"
	//"github.com/jroimartin/gocui"
)

type config struct {
	Game string
	Host string
	Port int
}

var (
	send    = make(chan string)
	receive = make(chan string)
)

func main() {
	g := flag.String("g", "Discworld", "Set name of game to connect to.")
	h := flag.String("h", "discworld.atuin.net", "Set host to connect to.")
	p := flag.Int("p", 4242, "Set port to connect to.")
	flag.Parse()

	cfg := &config{*g, *h, *p}

	c, err := gotelnet.Dial(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	read, write := bufio.NewReader(c), bufio.NewWriter(c)

	go func() {
		for {
			l := input.ReadLine()
			write.WriteString(l + "\n")
		}
	}()

	for {
		str, _ := read.ReadString('\n')
		fmt.Print(str)
	}
}
