package main

import (
	"bufio"
	"io"
)

func recvChan(c io.Reader) chan string {
	r := make(chan string)

	go func() {
		bufInput := bufio.NewReader(c)
		for {
			str, _ := bufInput.ReadString('\n')
			r <- str
		}
	}()

	return r
}
