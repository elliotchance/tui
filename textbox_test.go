package tui

import (
	"testing"
)

var textboxTests = map[string]windowTest{
	"less text than length": {
		3, 5,
		`.....
		 .....
		 .....`,
		`Hi~~~
		 ~~~~~
		 ~~~~~`,
		func(w *Window) {
			textBox := w.View().AddTextBox("Hi")
			textBox.Size().SetHeight(1)
			textBox.Size().SetWidth(4)
		},
		map[byte]Color{},
	},
	"exact text length": {
		3, 5,
		`.....
		 .....
		 .....`,
		`Foo~~
		 ~~~~~
		 ~~~~~`,
		func(w *Window) {
			textBox := w.View().AddTextBox("Foo")
			textBox.Size().SetHeight(1)
			textBox.Size().SetWidth(3)
		},
		map[byte]Color{},
	},
	"hide overflow text": {
		3, 5,
		`.....
		 .....
		 .....`,
		`Hell~
		 ~~~~~
		 ~~~~~`,
		func(w *Window) {
			textBox := w.View().AddTextBox("Hello World")
			textBox.Size().SetHeight(1)
			textBox.Size().SetWidth(4)
		},
		map[byte]Color{},
	},
	"default width is 100 percent": {
		3, 5,
		`.....
		 .....
		 .....`,
		`Hi~~~
		 ~~~~~
		 ~~~~~`,
		func(w *Window) {
			textBox := w.View().AddTextBox("Hi")
			textBox.Size().SetHeight(1)
		},
		map[byte]Color{},
	},
	"default width overflow": {
		3, 5,
		`.....
		 .....
		 .....`,
		`Hello
		 ~~~~~
		 ~~~~~`,
		func(w *Window) {
			textBox := w.View().AddTextBox("Hello World")
			textBox.Size().SetHeight(1)
		},
		map[byte]Color{},
	},
	"wrap lines": {
		3, 5,
		`.....
		 .....
		 .....`,
		`Foo~A
		 Bar~~
		 Bazzi`,
		func(w *Window) {
			w.View().AddTextBox("Foo A Bar Bazzing")
		},
		map[byte]Color{},
	},
}

func TestTextBox(t *testing.T) {
	runWindowTests(t, textboxTests)
}
