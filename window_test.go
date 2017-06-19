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

func addBorder(p [][]Pixel) (q [][]Pixel) {
	height := len(p)
	width := len(p[0])
	q = NewPixels(height+2, width+2, NoColor)

	// Create box
	q[0][0].Character = '+'
	q[height+1][0].Character = '+'
	q[0][width+1].Character = '+'
	q[height+1][width+1].Character = '+'

	for i := 0; i < width; i++ {
		q[0][i+1].Character = '-'
		q[height+1][i+1].Character = '-'
	}

	for i := 0; i < height; i++ {
		q[i+1][0].Character = '|'
		q[i+1][width+1].Character = '|'
	}

	// Fill in contents
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			q[i+1][j+1] = p[i][j]
		}
	}

	return
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
				// Go will remove spaces from strings which may affect how it
				// looks in the diff. WE will render a square around it.
				expectedPixels = addBorder(expectedPixels)
				actualPixels = addBorder(actualPixels)

				t.Errorf("expected:\n%s", Display(expectedPixels))
				t.Errorf("got:\n%s", Display(actualPixels))
			}
		})
	}
}

func TestWindow(t *testing.T) {
	runWindowTests(t, windowTests)
}
