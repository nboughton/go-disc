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

	if v, err := g.SetView(vLeftSide, -1, -1, int(0.2*float32(maxX)), maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Map"
	}

	if v, err := g.SetView(vMain, int(0.2*float32(maxX)), -1, maxX, maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Set some view paramters
		v.Title = "Main"
		v.Autoscroll = true
		v.Wrap = true

		// Opening Message
		fmt.Fprintln(v, "Welcome to go-disc")

		go func() {
			for {
				str, _ := dataIn.ReadString('\n')
				fmt.Fprint(v, str+"\n")
			}
		}()
	}

	if v, err := g.SetView(vInput, -1, maxY-5, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// View settings
		v.Title = "Input"
		v.Editable = true
		v.Wrap = true

		// Set focus on input
		if _, err := g.SetCurrentView(vInput); err != nil {
			return err
		}
	}

	return nil
}

func uiKeybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, uiQuit); err != nil {
		return err
	}

	// Submit a line
	if err := g.SetKeybinding(vInput, gocui.KeyEnter, gocui.ModNone, send); err != nil {
		return err
	}

	return nil
}

func uiQuit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
