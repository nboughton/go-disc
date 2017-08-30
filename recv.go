package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

var recv = make(chan string)

func recvHandler(v *gocui.View, r chan string) {
	for d := range r {
		if _, err := fmt.Fprint(v, d); err != nil {
			log.Println("RECV ERR:", err)
		}
	}
}
