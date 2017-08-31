package main

import (
	"github.com/jroimartin/gocui"
)

const (
	vMain     = "mainview"
	vLeftSide = "leftsideview"
	vInput    = "inputview"
)

func uiLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	/*
		if v, err := g.SetView(vLeftSide, 0, 0, int(0.2*float32(maxX)), maxY); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}

			v.Title = "Map"
		}
	*/

	//if v, err := g.SetView(vMain, int(0.2*float32(maxX)), 0, maxX, maxY); err != nil {
	if v, err := g.SetView(vMain, 0, 0, maxX-1, maxY-4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Set some view paramters
		v.Title = cfg.Session
		//v.Autoscroll = true
		//v.Wrap = true
		v.Frame = true
	}

	if v, err := g.SetView(vInput, 0, maxY-3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// View settings
		v.Title = "Input"
		v.Editable = true
		v.Wrap = true
		v.Highlight = true

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
