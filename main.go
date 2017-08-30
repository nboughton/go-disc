package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
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

	go receiveHandler(receive)
	go sendHandler(c)

	bufInput := bufio.NewReader(c)
	for {
		str, _ := bufInput.ReadString('\n')
		receive <- str
	}
}

func sendHandler(w io.Writer) {
	for {
		str := input.ReadLine()
		if _, err := w.Write([]byte(str + "\n")); err != nil {
			log.Println("SEND ERR:", err)
		}
	}
}

func receiveHandler(r chan string) {
	for d := range r {
		if _, err := fmt.Print(d); err != nil {
			log.Println("RECV ERR:", err)
		}
	}
}
