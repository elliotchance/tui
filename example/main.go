package main

import "github.com/elliotchance/tui"

func main() {
	window := tui.MainWindow()
	window.View().SetBackgroundColor(tui.NewColorRGB(0.2, 0.2, 0.2))

	pixels := window.View().Render()
	tui.Display(pixels)
}
