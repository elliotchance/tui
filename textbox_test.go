package tui

import (
	"testing"
)

var textboxTests = map[string]windowTest{
	"less text than length": {
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
