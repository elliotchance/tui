package tui

import (
	"reflect"
	"strings"
	"testing"
)

type windowTest struct {
	expectedColors string
	expectedText   string
	setup          func(*Window)
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
		func(w *Window) {
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
		func(w *Window) {
			w.SetBackgroundColor(Red)
		},
		map[byte]Color{
			'#': Red,
		},
	},
}

func stripSpace(s string) string {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\t", "", -1)

	return s
}

func runWindowTests(t *testing.T, tests map[string]windowTest) {
	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
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
				t.Errorf("expected:\n%s", Display(expectedPixels))
				t.Errorf("got:\n%s", Display(actualPixels))
			}
		})
	}
}

func TestWindow(t *testing.T) {
	runWindowTests(t, windowTests)
}
