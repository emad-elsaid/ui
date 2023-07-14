package ui

func ExampleChart() {
	exampleLoop(
		Rows(
			Flexed(1, Chart([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 100)),
			Flexed(1, Chart([]int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 9, 1, 5, 1}, 100)),
		),
	)
	// Output:
}
