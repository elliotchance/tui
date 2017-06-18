package tui

import (
	"testing"
)

var textboxTests = map[string]windowTest{
	"textbox less text": {
		`.....
		 .....
		 .....`,
		`Hi~~~
		 ~~~~~
		 ~~~~~`,
		func(w Window) {
			textBox := w.View().AddTextBox("Hi")
			textBox.Size().SetWidth(4)
		},
		map[byte]Color{},
	},
	"textbox exact text": {
		`.....
		 .....
		 .....`,
		`Foo~~
		 ~~~~~
		 ~~~~~`,
		func(w Window) {
			textBox := w.View().AddTextBox("Foo")
			textBox.Size().SetWidth(3)
		},
		map[byte]Color{},
	},
	"textbox overflow text": {
		`.....
		 .....
		 .....`,
		`Hell~
		 ~~~~~
		 ~~~~~`,
		func(w Window) {
			textBox := w.View().AddTextBox("Hello World")
			textBox.Size().SetWidth(4)
		},
		map[byte]Color{},
	},
	"textbox default width": {
		`.....
		 .....
		 .....`,
		`Hi~~~
		 ~~~~~
		 ~~~~~`,
		func(w Window) {
			w.View().AddTextBox("Hi")
		},
		map[byte]Color{},
	},
	"textbox default width overflow": {
		`.....
		 .....
		 .....`,
		`Hello
		 ~~~~~
		 ~~~~~`,
		func(w Window) {
			w.View().AddTextBox("Hello World")
		},
		map[byte]Color{},
	},
}

func TestTextBox(t *testing.T) {
	runWindowTests(t, textboxTests)
}
