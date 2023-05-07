package ui

import (
	"image/color"
)

func ExampleRoundedCorners() {
	box := Inset1(
		RoundedCorners(
			Background(
				STRAWBERRY_100,
				Inset1(
					Text("Hello world"),
				),
			),
		),
	)

	exampleLoop(
		Rows(
			Flexed(1, box),
			Flexed(1, box),
			Flexed(1, box),
		),
	)
	// Output:
}

func ExampleBorder() {
	box := Inset1(
		Border(
			Inset1(
				Text("Normal Border"),
			),
		),
	)

	boxActive := Inset1(
		BorderActive(
			Inset1(
				Text("Active Border"),
			),
		),
	)

	exampleLoop(
		Rows(
			Flexed(1, box),
			Flexed(1, boxActive),
		),
	)
	// Output:
}

func ExampleBackground() {
	colors := []color.NRGBA{
		STRAWBERRY_100, STRAWBERRY_300, STRAWBERRY_500, STRAWBERRY_700, STRAWBERRY_900,
		ORANGE_100, ORANGE_300, ORANGE_500, ORANGE_700, ORANGE_900,
		BANANA_100, BANANA_300, BANANA_500, BANANA_700, BANANA_900,
		LIME_100, LIME_300, LIME_500, LIME_700, LIME_900,
		MINT_100, MINT_300, MINT_500, MINT_700, MINT_900,
		BLUEBERRY_100, BLUEBERRY_300, BLUEBERRY_500, BLUEBERRY_700, BLUEBERRY_900,
		BUBBLEGUM_100, BUBBLEGUM_300, BUBBLEGUM_500, BUBBLEGUM_700, BUBBLEGUM_900,
		GRAPE_100, GRAPE_300, GRAPE_500, GRAPE_700, GRAPE_900,
		COCOA_100, COCOA_300, COCOA_500, COCOA_700, COCOA_900,
		SILVER_100, SILVER_300, SILVER_500, SILVER_700, SILVER_900,
		SLATE_100, SLATE_300, SLATE_500, SLATE_700, SLATE_900,
		BLACK_100, BLACK_300, BLACK_500, BLACK_700, BLACK_900,
	}

	l := NewVerticalList()
	w := func(c C) D {
		return l.Layout(c, len(colors), func(c C, i int) D {
			return Background(colors[i], Text(""))(c)
		})
	}

	exampleLoop(w)
	// Output:
}

func ExampleHR() {
	exampleLoop(
		Inset1(
			Rows(
				Flexed(1, Text("Line 1")),
				Flexed(1, HR(1)),
				Flexed(1, Text("Line 2")),
				Flexed(1, HR(2)),
				Flexed(1, Text("Line 3")),
			),
		),
	)
	// Output:
}

func ExampleVR() {
	exampleLoop(
		Inset1(
			Columns(
				Flexed(1, Text("Line 1")),
				Flexed(1, VR(1)),
				Flexed(1, Text("Line 2")),
				Flexed(1, VR(2)),
				Flexed(1, Text("Line 3")),
			),
		),
	)
	// Output:
}
