package main

import (
	"io"
	"log"
	//"github.com/nboughton/go-utils/input"
)

var send = make(chan string)

func sendHandler(s chan string, w io.Writer) {
	for str := range s {
		//str := input.ReadLine()
		if _, err := w.Write([]byte(str + "\n")); err != nil {
			log.Println("SEND ERR:", err)
		}
	}
}
