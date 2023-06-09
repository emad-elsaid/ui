package ui

import (
	"gioui.org/font"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/widget"
)

func Text(s string) W {
	return func(c C) D {
		tl := widget.Label{Alignment: Theme.TextAlignment, MaxLines: Theme.MaxLines}
		paint.ColorOp{Color: Theme.TextColor}.Add(c.Ops)
		return tl.Layout(c, Theme.FontFamily, font.Font{Weight: Theme.FontWeight}, Theme.FontSize, s, op.CallOp{})
	}
}

var OneLine = MaxLines(1)

func Label(s string) W { return OneLine(Text(s)) }

var Bold = FontWeight(font.Bold)
