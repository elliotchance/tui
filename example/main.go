package main

import (
	"github.com/elliotchance/tui"
	"github.com/nsf/termbox-go"
)

func main() {
	window := tui.MainWindow()
	view := window.View()
	splitView := view.AddSplitView(20)

	nav := splitView.LeftView().AddListView()
	nav.SetItems([]string{
		"Blue background",
		"Green background",
	})

	splitView.RightView().SetBackgroundColor(tui.Color(tui.Blue))
	splitView.RightView().AddTextBox("Press the down arrow...")

	window.Start(func(ev termbox.Event) bool {
		switch ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyArrowDown {
				nav.SetSelectedIndex(1)
				splitView.RightView().SetBackgroundColor(tui.Color(tui.Red))
			}

			if ev.Key == termbox.KeyArrowUp {
				nav.SetSelectedIndex(0)
				splitView.RightView().SetBackgroundColor(tui.Color(tui.Blue))
			}

			if ev.Key == termbox.KeyCtrlC {
				return false
			}
		}

		return true
	})
}
