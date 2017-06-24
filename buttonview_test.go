package tui

import (
	"testing"
)

var buttonviewTests = map[string]windowTest{
	"no buttons": {
		3, 5,
		`.....
		 .....
		 .....`,
		`~~~~~
		 ~~~~~
		 ~~~~~`,
		func(w *Window) {
			w.View().AddButtonView([]Button{})
		},
		map[byte]Color{},
	},
	"one button": {
		3, 5,
		`####.
		 ####.
		 ####.`,
		`~~~~~
		 ~OK~~
		 ~~~~~`,
		func(w *Window) {
			w.View().AddButtonView([]Button{{"OK"}})
		},
		map[byte]Color{
			'#': Blue,
		},
	},
	"two buttons": {
		3, 13,
		`########.####
		 ########.####
		 ########.####`,
		`~~~~~~~~~~~~~
		 ~Cancel~~~OK~
		 ~~~~~~~~~~~~~`,
		func(w *Window) {
			w.View().AddButtonView([]Button{{"Cancel"}, {"OK"}})
		},
		map[byte]Color{
			'#': Blue,
		},
	},
}

func TestButtonView(t *testing.T) {
	runWindowTests(t, buttonviewTests)
}
