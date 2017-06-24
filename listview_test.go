package tui

import (
	"testing"
)

var listviewTests = map[string]windowTest{
	"no items": {
		3, 5,
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
	"one item": {
		3, 5,
		`#####
		 .....
		 .....`,
		`Foo~~
		 ~~~~~
		 ~~~~~`,
		func(w *Window) {
			listView := w.View().AddListView()
			listView.SetItems([]string{"Foo"})
		},
		map[byte]Color{
			'#': Blue,
		},
	},
	"three items": {
		3, 5,
		`#####
		 .....
		 .....`,
		`Foo~~
		 Bar~~
		 Baz~~`,
		func(w *Window) {
			listView := w.View().AddListView()
			listView.SetItems([]string{"Foo", "Bar", "Baz"})
		},
		map[byte]Color{
			'#': Blue,
		},
	},
	"item truncated": {
		3, 5,
		`#####
		 .....
		 .....`,
		`Foo_B
		 ~~~~~
		 ~~~~~`,
		func(w *Window) {
			listView := w.View().AddListView()
			listView.SetItems([]string{"Foo Bar"})
		},
		map[byte]Color{
			'#': Blue,
		},
	},
	"too many items": {
		3, 5,
		`#####
		 .....
		 .....`,
		`Foo~~
		 Bar~~
		 Baz~~`,
		func(w *Window) {
			listView := w.View().AddListView()
			listView.SetItems([]string{"Foo", "Bar", "Baz", "Qux"})
		},
		map[byte]Color{
			'#': Blue,
		},
	},
	"background color": {
		3, 5,
		`+++++
		 #####
		 #####`,
		`Foo~~
		 Bar~~
		 Baz~~`,
		func(w *Window) {
			listView := w.View().AddListView()
			listView.SetItems([]string{"Foo", "Bar", "Baz"})
			listView.SetBackgroundColor(Green)
		},
		map[byte]Color{
			'#': Green,
			'+': Blue,
		},
	},
}

func TestListView(t *testing.T) {
	runWindowTests(t, listviewTests)
}
