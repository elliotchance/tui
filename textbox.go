package tui

import (
	"strings"
	"github.com/nsf/termbox-go"
)

type TextBox struct {
	text            string
	backgroundColor Color
	size            *Size
}

func (v *View) AddTextBox(text string) *TextBox {
	textBox := &TextBox{
		text:            text,
		backgroundColor: NoColor,
		size:            newMutableSize(v.Size().AbsoluteLeft(), v.Size().AbsoluteTop(), v.Size().Height(), v.Size().Width()),
	}

	v.child = textBox

	return textBox
}

func (v *TextBox) SetText(text string) {
	v.text = text
}

func (v *TextBox) Text() string {
	return v.text
}

func (v *TextBox) SetBackgroundColor(c Color) {
	v.backgroundColor = c
}

func (v *TextBox) BackgroundColor() Color {
	return v.backgroundColor
}

func (v *TextBox) Size() MutableSizer {
	return v.size
}

func (v *TextBox) Render() (rows [][]Pixel) {
	rows = NewPixels(v.Size().Height(), v.Size().Width(), v.backgroundColor)

	// Split the text into words as we might need to wrap the text over multiple
	// lines.
	words := strings.Split(v.text, " ")

	lineNumber := 0
	columnNumber := 0
	for _, word := range words {
		nextWordLength := len(word)
		if columnNumber > 0 {
			// Make room for the space if its not the first word in the line.
			nextWordLength++
		}

		// If the word does not fit on this line we need to end the line now.
		if columnNumber+nextWordLength > v.Size().Width() {
			// If there are no more available lines we should finish here.
			if lineNumber >= len(rows)-1 {
				// Always print any remaining characters in the left over space
				// to indicate that there is more text. This is also important
				// when placing a single word in a textbox that is longer than
				// the textbox.
				for _, c := range word {
					if columnNumber >= v.Size().Width() {
						return
					}

					rows[lineNumber][columnNumber].Character = c
					columnNumber++
				}

				break
			}

			// Move the cursor to the start of the next line.
			lineNumber++
			columnNumber = 0
		}

		// If this is not the first word on the line we need to add a space
		// before the next word.
		if columnNumber > 0 {
			rows[lineNumber][columnNumber].Character = ' '
			columnNumber++
		}

		for _, c := range word {
			// Make sure we never overflow.
			if columnNumber >= v.Size().Width() {
				return
			}

			rows[lineNumber][columnNumber].Character = c
			columnNumber++
		}
	}

	return
}

func (v *TextBox) setContainerSize(left, top, height, width int) {
	v.Size().setContainerSize(left, top, height, width)
}

func (v *TextBox) getViewForPosition(x, y int) Renderer {
	return v
}

func (v *TextBox) handleEvent(e termbox.Event) {
}
