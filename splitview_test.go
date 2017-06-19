package tui

import (
	"testing"
)

var splitviewTests = map[string]windowTest{
	"default width": {
		`##+++
		 ##+++
		 ##+++`,
		`~~~~~
		 ~~~~~
		 ~~~~~`,
		func(w *Window) {
			splitView := w.View().AddSplitView()
			splitView.LeftView().SetBackgroundColor(Red)
			splitView.RightView().SetBackgroundColor(Green)
		},
		map[byte]Color{
			'#': Red,
			'+': Green,
		},
	},
}

func TestSplitView(t *testing.T) {
	runWindowTests(t, splitviewTests)
}
