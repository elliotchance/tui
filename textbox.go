package tui

import (
	"strings"
)

type TextBox interface {
	SetText(string)
	Text() string
	Size() MutableSize
	Render() [][]Pixel
}

type textBox struct {
	text            string
	backgroundColor Color
	size            MutableSize
}

func (v *view) AddTextBox(text string) TextBox {
	textBox := &textBox{
		text:            text,
		backgroundColor: NoColor,
		size:            newMutableSize(v.Size().Height(), v.Size().Width()),
	}

	v.child = textBox

	return textBox
}

func (v *textBox) SetText(text string) {
	v.text = text
}

func (v *textBox) Text() string {
	return v.text
}

func (v *textBox) SetBackgroundColor(c Color) {
	v.backgroundColor = c
}

func (v *textBox) BackgroundColor() Color {
	return v.backgroundColor
}

func (v *textBox) Size() MutableSize {
	return v.size
}

func (v *textBox) Render() (rows [][]Pixel) {
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
