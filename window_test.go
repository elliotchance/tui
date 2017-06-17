package tui

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type windowTest struct {
	expectedColors string
	expectedText   string
	setup          func(Window)
	pixelMap       map[byte]Color
}

var windowTests = map[string]windowTest{
	"blank window": {
		`.....
		 .....
		 .....`,
		`~~~~~
		 ~~~~~
		 ~~~~~`,
		func(w Window) {
		},
		map[byte]Color{},
	},
	"red window": {
		`#####
		 #####
		 #####`,
		`~~~~~
		 ~~~~~
		 ~~~~~`,
		func(w Window) {
			w.SetBackgroundColor(Red)
		},
		map[byte]Color{
			'#': Red,
		},
	},
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

func stripSpace(s string) string {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\t", "", -1)

	return s
}

func TestWindow(t *testing.T) {
	for testName, test := range windowTests {
		height, width := 3, 5
		window := newWindow(height, width)

		test.setup(window)

		expectedColors := stripSpace(test.expectedColors)
		expectedText := stripSpace(test.expectedText)

		expectedPixels := NewPixels(height, width, NoColor)
		for row := 0; row < len(expectedPixels); row++ {
			for col := 0; col < len(expectedPixels[0]); col++ {
				pos := len(expectedPixels[0])*row + col

				char := expectedText[pos]
				if char == '~' {
					char = ' '
				}

				color := test.pixelMap[expectedColors[pos]]
				if expectedColors[pos] == '.' {
					color = NoColor
				}

				expectedPixels[row][col].Character = rune(char)
				expectedPixels[row][col].BackgroundColor = color
			}
		}

		actualPixels := window.Render()

		if !reflect.DeepEqual(actualPixels, expectedPixels) {
			fmt.Printf("%s expected:\n", testName)
			Display(expectedPixels)
			fmt.Printf("%s got:\n", testName)
			Display(actualPixels)

			t.Fail()
		}
	}
}
