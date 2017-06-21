package tui

import (
	"testing"
)

var splitviewTests = map[string]windowTest{
	"flexible": {
		`##+++
		 ##+++
		 ##+++`,
		`~~~~~
		 ~~~~~
		 ~~~~~`,
		func(w *Window) {
			splitView := w.View().AddFlexibleSplitView(0.5)
			splitView.LeftView().SetBackgroundColor(Red)
			splitView.RightView().SetBackgroundColor(Green)
		},
		map[byte]Color{
			'#': Red,
			'+': Green,
		},
	},
	"fixed": {
		`#++++
		 #++++
		 #++++`,
		`~~~~~
		 ~~~~~
		 ~~~~~`,
		func(w *Window) {
			splitView := w.View().AddSplitView(1)
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
