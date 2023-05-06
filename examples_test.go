package ui

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
)

func exampleLoop(w W) {
	win := app.NewWindow()
	var ops op.Ops

	for e := range win.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			c := layout.NewContext(&ops, e)
			w(c)
			e.Frame(c.Ops)
		case system.DestroyEvent:
			return
		}
	}
}
