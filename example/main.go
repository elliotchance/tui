package main

import "github.com/elliotchance/tui"

func main() {
	window := tui.MainWindow()
	view := window.View()
	splitView := view.AddSplitView()

	nav := splitView.LeftView().AddListView()
	nav.SetItems([]string{
		"Press the down arrow...",
		"Blue background",
	})

	splitView.RightView().SetBackgroundColor(tui.Blue)

	pixels := window.Render()
	tui.Print(pixels)
}
