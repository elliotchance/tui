package tui

import (
	"testing"
)

var viewTests = map[string]windowTest{
	"green window view": {
		`#####
		 #####
		 #####`,
		`~~~~~
		 ~~~~~
		 ~~~~~`,
		func(w *Window) {
			w.View().SetBackgroundColor(Green)
		},
		map[byte]Color{
			'#': Green,
		},
	},
	"resize window view": {
		`###..
		 ###..
		 .....`,
		`~~~~~
		 ~~~~~
		 ~~~~~`,
		func(w *Window) {
			w.View().SetBackgroundColor(Green)
			w.View().Size().SetHeight(2)
			w.View().Size().SetWidth(3)
		},
		map[byte]Color{
			'#': Green,
		},
	},
	"flexible window view": {
		`####.
		 ####.
		 .....`,
		`~~~~~
		 ~~~~~
		 ~~~~~`,
		func(w *Window) {
			w.View().SetBackgroundColor(Green)
			w.View().Size().SetFlexibleHeight(0.7)
			w.View().Size().SetFlexibleWidth(0.8)
		},
		map[byte]Color{
			'#': Green,
		},
	},
}

func TestView(t *testing.T) {
	runWindowTests(t, viewTests)
}
