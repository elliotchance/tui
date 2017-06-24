package tui

import "github.com/nsf/termbox-go"

type Button struct {
	text string
}

type ButtonView struct {
	backgroundColor Color
	size            *Size
	buttons         []Button
}

func (v *View) AddButtonView(buttons []Button) *ButtonView {
	buttonView := &ButtonView{
		backgroundColor: NoColor,
		size:            newMutableSize(v.Size().AbsoluteLeft(), v.Size().AbsoluteTop(), v.Size().Height(), v.Size().Width()),
		buttons:         buttons,
	}

	v.child = buttonView

	return buttonView
}

func (v *ButtonView) SetBackgroundColor(c Color) {
	v.backgroundColor = c
}

func (v *ButtonView) BackgroundColor() Color {
	return v.backgroundColor
}

func (v *ButtonView) Size() MutableSizer {
	return v.size
}

func (v *ButtonView) Render() [][]Pixel {
	width := v.Size().Width()
	height := v.Size().Height()
	rows := NewPixels(height, width, v.backgroundColor)

	// The buttons will always sit at the bottom. It may make more sense to have
	// them in the middle.

	// TODO: The view size must be greater than 5 height or things start to go
	// wrong.

	i := 0
	for _, button := range v.buttons {
		rows[len(rows)-3][i].BackgroundColor = Blue
		rows[len(rows)-2][i].BackgroundColor = Blue
		rows[len(rows)-1][i].BackgroundColor = Blue
		i++

		for _, c := range button.text {
			// Line above
			rows[len(rows)-3][i].BackgroundColor = Blue

			// Button line
			rows[len(rows)-2][i].Character = c
			rows[len(rows)-2][i].BackgroundColor = Blue

			// Line below
			rows[len(rows)-1][i].BackgroundColor = Blue

			i++
		}

		rows[len(rows)-3][i].BackgroundColor = Blue
		rows[len(rows)-2][i].BackgroundColor = Blue
		rows[len(rows)-1][i].BackgroundColor = Blue

		i += 2
	}

	return rows
}

func (v *ButtonView) setContainerSize(left, top, height, width int) {
	v.Size().setContainerSize(left, top, height, width)
}

func (v *ButtonView) getViewForPosition(x, y int) Renderer {
	return v
}

func (v *ButtonView) handleEvent(e termbox.Event) {
}
