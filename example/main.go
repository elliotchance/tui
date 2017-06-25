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
	nav.OnIndexChanged(func() {
		switch nav.SelectedIndex() {
		case 0:
			splitView.RightView().SetBackgroundColor(tui.Color(tui.Blue))
		case 1:
			splitView.RightView().SetBackgroundColor(tui.Color(tui.Red))
		case 3:
			modal := tui.NewModal(12, 50)
			modal.View().AddTextBox("Hello, I'm a modal.")

			window.ShowModal(modal)
		}
	})
	nav.SetItems([]string{
		"Blue background",
		"Red background",
		"Colors",
		"Modals",
	})

	window.Start()
}
