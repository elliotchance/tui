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
		func(w Window) {
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
		func(w Window) {
			w.View().SetBackgroundColor(Green)
			w.View().(ResizableView).SetHeight(2)
			w.View().(ResizableView).SetWidth(3)
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
		func(w Window) {
			w.View().SetBackgroundColor(Green)
			w.View().(ResizableView).SetFlexibleHeight(0.7)
			w.View().(ResizableView).SetFlexibleWidth(0.8)
		},
		map[byte]Color{
			'#': Green,
		},
	},
}

func TestView(t *testing.T) {
	runWindowTests(t, viewTests)
}
