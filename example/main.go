package main

import "github.com/elliotchance/tui"

func main() {
	window := tui.MainWindow()
	view := window.View()
	// view.SetBackgroundColor(tui.NewColorRGB(0.2, 0.2, 0.2))
	// view.SetBackgroundColor(tui.NoColor)

	pixels := view.Render()
	tui.Display(pixels)
}
