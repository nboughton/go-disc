package main

import (
	"github.com/nboughton/go-disc/complete"
	"github.com/nboughton/go-disc/history"
	"github.com/stesla/gotelnet"
)

type config struct {
	Session string
	Host    string
	Port    int
}

const (
	vMain     = "mainview"
	vLeftSide = "leftsideview"
	vInput    = "inputview"
	minX      = -1
	minY      = -1
)

var (
	cfg  *config
	conn gotelnet.Conn
	cmds = history.New()
	dict = complete.New()
)
