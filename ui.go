package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

const (
	vMain     = "mainview"
	vLeftSide = "leftsideview"
	vInput    = "inputview"
)

func uiLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if _, err := g.SetView(vLeftSide, -1, -1, int(0.2*float32(maxX)), maxY-5); err != nil &&
		err != gocui.ErrUnknownView {
		return err
	}

	if v, err := g.SetView(vMain, int(0.2*float32(maxX)), -1, maxX, maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Set some view paramters
		v.Autoscroll = true
		v.Wrap = true

		// Opening Message
		fmt.Fprintln(v, "Welcome to go-disc")
	}

	if _, err := g.SetView(vInput, -1, maxY-5, maxX, maxY); err != nil &&
		err != gocui.ErrUnknownView {
		return err
	}
	return nil
}

func uiQuit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
