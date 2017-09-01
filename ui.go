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

var (
	cmdBuffer []string
	cmdIdx    int
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
	if v, err := g.SetView(vMain, -1, -1, maxX, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Set some view paramters
		//v.Title = cfg.Session
		v.Autoscroll = true
		v.Wrap = true
	}

	if v, err := g.SetView(vInput, -1, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// View settings
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
	if err := g.SetKeybinding(vInput, gocui.KeyEnter, gocui.ModNone, input); err != nil {
		return err
	}

	// Scroll cmd buffer
	if err := g.SetKeybinding(vInput, gocui.KeyArrowUp, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			scrollCmdHistory(v, -1)
			return nil
		}); err != nil {
		return err
	}

	if err := g.SetKeybinding(vInput, gocui.KeyArrowDown, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			scrollCmdHistory(v, 1)
			return nil
		}); err != nil {
		return err
	}

	// Tab completion
	if err := g.SetKeybinding(vInput, gocui.KeyTab, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			x, y := v.Cursor()
			if x > 0 {
				x--
			}
			str, _ := v.Word(x, y)

			//vM, _ := g.View(vMain)
			//fmt.Fprintf(vM, "x: %v, y: %v : %v\n", x, y, str)

			tab, _ := dict.Tab(str)
			fmt.Fprintf(v, "%s", tab)
			return nil
		}); err != nil {
		return err
	}

	return nil
}

func scrollCmdHistory(v *gocui.View, dy int) {
	i := cmdIdx + dy
	switch {
	case i >= 0 && i < len(cmdBuffer):
		cmdIdx = i

		v.Clear()
		fmt.Fprintf(v, "%v", cmdBuffer[cmdIdx])
	case i == len(cmdBuffer):
		v.Clear()
		v.SetOrigin(0, 0)
	}
}

func uiQuit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
