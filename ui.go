package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
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

	if v, err := g.SetView(vTop, minX, minY, maxX, minY+6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Set view parameters
		textBox(v)
	}

	if v, err := g.SetView(vMain, minX, minY+6, maxX, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Set some view paramters
		textBox(v)
	}

	if v, err := g.SetView(vInput, minX, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// View settings
		v.Frame = false
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

	// home and end keys
	if err := g.SetKeybinding(vInput, gocui.KeyHome, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return v.SetCursor(0, 0)
		}); err != nil {
		return err
	}

	if err := g.SetKeybinding(vInput, gocui.KeyEnd, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			n := len(v.Buffer()) - 1
			if n < 0 {
				return nil
			}
			return v.SetCursor(n, 0)
		}); err != nil {
		return err
	}

	// Submit a line
	if err := g.SetKeybinding(vInput, gocui.KeyEnter, gocui.ModNone, input); err != nil {
		return err
	}

	// Scroll cmd buffer
	if err := g.SetKeybinding(vInput, gocui.KeyArrowUp, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return scrollCmdHistory(v, true)
		}); err != nil {
		return err
	}

	if err := g.SetKeybinding(vInput, gocui.KeyArrowDown, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return scrollCmdHistory(v, false)
		}); err != nil {
		return err
	}

	// Tab completion
	if err := g.SetKeybinding(vInput, gocui.KeyTab, gocui.ModNone, tabComplete); err != nil {
		return err
	}

	return nil
}

func tabComplete(g *gocui.Gui, v *gocui.View) error {
	// Get cursor coords
	x, y := v.Cursor()

	// We don't complete empty lines
	if x == 0 {
		return nil
	}

	// Extract substr of last word behind cursor point
	line, _ := v.Line(y)
	str := ""
	for i := x - 1; i >= 0; i-- {
		if line[i] == ' ' {
			str = line[i+1 : x]
			break
		} else if i == 0 {
			str = line[i:x]
			break
		}
	}

	// Return clear if we've come up empty handed
	if str == "" {
		return nil
	}

	// Attempt to get a tab
	tab, _ := dict.Tab(str)

	// Catch the current line, zero the content and print the line with the completion
	zeroLine(v)
	fmt.Fprintf(v, "%s%s", line[:x], tab)

	// Set cursor to its original position in case we need to tab again.
	v.SetCursor(x, y)
	return nil
}

func scrollCmdHistory(v *gocui.View, back bool) error {
	v.Clear()
	if back {
		fmt.Fprint(v, cmds.Prev())
	} else {
		fmt.Fprint(v, cmds.Next())
	}

	return nil
}

func uiQuit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func textBox(v *gocui.View) {
	//v.Frame = false
	v.Autoscroll = true
	v.Wrap = true
}

func zeroLine(v *gocui.View) {
	v.SetCursor(0, 0)
	v.SetOrigin(0, 0)
	v.Clear()
}
