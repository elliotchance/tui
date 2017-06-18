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
			textBox := NewTextBox("Hi")
			textBox.SetWidth(4)
			w.View().AddChild(textBox)
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
			textBox := NewTextBox("Foo")
			textBox.SetWidth(3)
			w.View().AddChild(textBox)
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
			textBox := NewTextBox("Hello World")
			textBox.SetWidth(4)
			w.View().AddChild(textBox)
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
			textBox := NewTextBox("Hi")
			w.View().AddChild(textBox)
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
			textBox := NewTextBox("Hello World")
			w.View().AddChild(textBox)
		},
		map[byte]Color{},
	},
}

func TestTextBox(t *testing.T) {
	runWindowTests(t, textboxTests)
}
