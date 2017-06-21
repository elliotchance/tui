package main

import "github.com/elliotchance/tui"

func main() {
	window := tui.MainWindow()
	view := window.View()
	splitView := view.AddSplitView(20)

	nav := splitView.LeftView().AddListView()
	nav.SetItems([]string{
		"Blue background",
	})

	splitView.RightView().SetBackgroundColor(tui.Blue)
	splitView.RightView().AddTextBox("Press the down arrow...")

	pixels := window.Render()
	tui.Print(pixels)
}
