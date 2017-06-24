package main

import (
	"github.com/elliotchance/tui"
)

func main() {
	window := tui.MainWindow()
	view := window.View()
	splitView := view.AddSplitView(20)

	boxView := splitView.LeftView().AddBoxView()
	boxView.SetTitle("Demos")

	nav := boxView.View().AddListView()
	nav.SetItems([]string{
		"Blue background",
		"Red background",
	})

	splitView.RightView().SetBackgroundColor(tui.Color(tui.Blue))
	splitView.RightView().AddTextBox("Press the down arrow...")

	window.Start()
}
