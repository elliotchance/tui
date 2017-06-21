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
	"one item": {
		`.....
		 .....
		 .....`,
		`Foo~~
		 ~~~~~
		 ~~~~~`,
		func(w *Window) {
			listView := w.View().AddListView()
			listView.SetItems([]string{"Foo"})
		},
		map[byte]Color{},
	},
	"three items": {
		`.....
		 .....
		 .....`,
		`Foo~~
		 Bar~~
		 Baz~~`,
		func(w *Window) {
			listView := w.View().AddListView()
			listView.SetItems([]string{"Foo", "Bar", "Baz"})
		},
		map[byte]Color{},
	},
	"item truncated": {
		`.....
		 .....
		 .....`,
		`Foo~B
		 ~~~~~
		 ~~~~~`,
		func(w *Window) {
			listView := w.View().AddListView()
			listView.SetItems([]string{"Foo Bar"})
		},
		map[byte]Color{},
	},
	"too many items": {
		`.....
		 .....
		 .....`,
		`Foo~~
		 Bar~~
		 Baz~~`,
		func(w *Window) {
			listView := w.View().AddListView()
			listView.SetItems([]string{"Foo", "Bar", "Baz", "Qux"})
		},
		map[byte]Color{},
	},
	"background color": {
		`#####
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
		},
	},
}

func TestListView(t *testing.T) {
	runWindowTests(t, listviewTests)
}
