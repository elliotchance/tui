package tui

import (
	"testing"
)

var listviewTests = map[string]windowTest{
	"no items": {
		`.....
		 .....
		 .....`,
		`~~~~~
		 ~~~~~
		 ~~~~~`,
		func(w *Window) {
			w.View().AddListView()
		},
		map[byte]Color{},
	},
}

func TestListView(t *testing.T) {
	runWindowTests(t, listviewTests)
}
