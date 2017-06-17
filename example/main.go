package main

import "github.com/elliotchance/tui"

func main() {
	window := tui.MainWindow()
	view := window.View()
	window.SetBackgroundColor(tui.NewColorRGB(0.2, 0.2, 0.2))

	view.SetBackgroundColor(tui.Blue)
	// view.SetWidth(20)
	// view.SetHeight(10)
	view.SetFlexibleWidth(0.5)
	view.SetFlexibleHeight(0.2)

	pixels := window.Render()
	tui.Display(pixels)
}
