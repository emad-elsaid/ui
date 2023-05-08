package ui

import (
	"image"
	"image/color"

	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
)

func RoundedCorners(w W) W {
	return func(c C) D {
		macro := op.Record(c.Ops)
		d := w(c)
		macroOp := macro.Stop()

		defer clip.UniformRRect(image.Rect(0, 0, d.Size.X, d.Size.Y), Theme.BorderRadius).Push(c.Ops).Pop()
		macroOp.Add(c.Ops)
		return d
	}
}

func Border(w W) W {
	return func(c C) D {
		return widget.Border{Color: Theme.BorderColor, Width: BorderSize, CornerRadius: 0}.Layout(c, w)
	}
}

func BorderActive(w W) W {
	return func(c C) D {
		return widget.Border{Color: Theme.ActiveBorderColor, Width: BorderSize, CornerRadius: 0}.Layout(c, w)
	}
}

func Background(background color.NRGBA, w W) W {
	return func(c C) D {
		macro := op.Record(c.Ops)
		d := w(c)
		path := macro.Stop()

		cl := clip.Rect{Max: d.Size}.Push(c.Ops)
		paint.Fill(c.Ops, background)
		cl.Pop()

		path.Add(c.Ops)
		return d
	}
}

func HR(sz int) W {
	return func(c C) D {
		cl := clip.Path{}
		cl.Begin(c.Ops)
		cl.MoveTo(Pt(0, 0))
		cl.Line(Pt(float32(c.Constraints.Max.X), 0))

		defer clip.Stroke{
			Path:  cl.End(),
			Width: float32(sz),
		}.Op().Push(c.Ops).Pop()

		paint.Fill(c.Ops, Theme.BorderColor)

		return D{Size: P{c.Constraints.Min.X, sz}}
	}
}

func VR(sz int) W {
	return func(c C) D {
		cl := clip.Path{}
		cl.Begin(c.Ops)
		cl.MoveTo(Pt(0, 0))
		cl.Line(Pt(0, float32(c.Constraints.Max.Y)))

		defer clip.Stroke{
			Path:  cl.End(),
			Width: float32(sz),
		}.Op().Push(c.Ops).Pop()

		paint.Fill(c.Ops, Theme.BorderColor)

		return D{Size: P{sz, c.Constraints.Min.Y}}
	}
}
