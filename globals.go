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

const (
	vMain     = "mainview"
	vLeftSide = "leftsideview"
	vInput    = "inputview"
	minX      = -1
	minY      = -1
)

var (
	conn           gotelnet.Conn
	cfg            *config
	cmdBuffer      []string
	cmdIdx         int
	logToCmdBuffer bool
	dict           = tc.New()
)
