package main

import (
	tc "github.com/nboughton/go-disc/complete"
	"github.com/stesla/gotelnet"
)

type config struct {
	Session string
	Host    string
	Port    int
}

type cmdHistory struct {
	buffer []string
	idx    int
	log    bool
}

const (
	vMain     = "mainview"
	vLeftSide = "leftsideview"
	vInput    = "inputview"
	minX      = -1
	minY      = -1
)

var (
	conn gotelnet.Conn
	cfg  *config
	cmds cmdHistory
	dict = tc.New()
)
