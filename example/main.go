package main

import (
	"github.com/elliotchance/tui"
)

func main() {
	window := tui.MainWindow()
	pixels := window.Render()
	tui.Display(pixels)
}
