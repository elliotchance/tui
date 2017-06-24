package tui

import (
	"testing"
)

var boxviewTests = map[string]windowTest{
	"with title": {
		3, 10,
		`..........
		 ..........
		 ..........`,
		`┌~Hi~────┐
		 │~~~~~~~~│
		 └────────┘`,
		func(w *Window) {
			boxView := w.View().AddBoxView()
			boxView.SetTitle("Hi")
		},
		map[byte]Color{},
	},
	"without title": {
		3, 10,
		`..........
		 ..........
		 ..........`,
		`┌────────┐
		 │~~~~~~~~│
		 └────────┘`,
		func(w *Window) {
			w.View().AddBoxView()
		},
		map[byte]Color{},
	},
	"background color": {
		3, 10,
		`##########
		 ##########
		 ##########`,
		`┌────────┐
		 │~~~~~~~~│
		 └────────┘`,
		func(w *Window) {
			boxView := w.View().AddBoxView()
			boxView.SetBackgroundColor(Green)
		},
		map[byte]Color{
			'#': Green,
		},
	},
}

func TestBoxView(t *testing.T) {
	runWindowTests(t, boxviewTests)
}
