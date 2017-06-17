package tui

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type windowTest struct {
	expected string
	setup    func(Window)
	pixelMap map[byte]Color
}

var windowTests = map[string]windowTest{
	"blank window": {
		`.....
		 .....
		 .....`,
		func(w Window) {
		},
		map[byte]Color{
			'.': NoColor,
		},
	},
	"red window": {
		`.....
		 .....
		 .....`,
		func(w Window) {
			w.SetBackgroundColor(Red)
		},
		map[byte]Color{
			'.': Red,
		},
	},
	"green window view": {
		`.....
		 .....
		 .....`,
		func(w Window) {
			w.View().SetBackgroundColor(Green)
		},
		map[byte]Color{
			'.': Green,
		},
	},
	"resize window view": {
		`###..
		 ###..
		 .....`,
		func(w Window) {
			w.View().SetBackgroundColor(Green)
			w.View().(ResizableView).SetHeight(2)
			w.View().(ResizableView).SetWidth(3)
		},
		map[byte]Color{
			'.': NoColor,
			'#': Green,
		},
	},
	"flexible window view": {
		`####.
		 ####.
		 .....`,
		func(w Window) {
			w.View().SetBackgroundColor(Green)
			w.View().(ResizableView).SetFlexibleHeight(0.7)
			w.View().(ResizableView).SetFlexibleWidth(0.8)
		},
		map[byte]Color{
			'.': NoColor,
			'#': Green,
		},
	},
}

func TestWindow(t *testing.T) {
	for testName, test := range windowTests {
		height, width := 3, 5
		window := newWindow(height, width)

		test.setup(window)

		expected := test.expected
		expected = strings.Replace(expected, " ", "", -1)
		expected = strings.Replace(expected, "\n", "", -1)
		expected = strings.Replace(expected, "\t", "", -1)

		expectedPixels := NewPixels(height, width, NoColor)
		actualPixels := window.Render()

		for row := 0; row < len(expectedPixels); row++ {
			for col := 0; col < len(expectedPixels[0]); col++ {
				pos := len(expectedPixels[0])*row + col
				expectedPixels[row][col].BackgroundColor = test.pixelMap[expected[pos]]
			}
		}

		if !reflect.DeepEqual(actualPixels, expectedPixels) {
			fmt.Printf("%s expected:\n", testName)
			Display(expectedPixels)
			fmt.Printf("%s got:\n", testName)
			Display(actualPixels)
			t.Fail()
		}
	}
}
