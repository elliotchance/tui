package tui

type TextBox interface {
	View

	SetText(string)
	Text() string
	SetWidth(int)
}

type textBox struct {
	text            string
	width           int
	backgroundColor Color
	containerWidth  int
}

func NewTextBox(text string) TextBox {
	return &textBox{
		text:            text,
		backgroundColor: NoColor,
		width:           -1,
	}
}

func (v *textBox) SetText(text string) {
	v.text = text
}

func (v *textBox) Text() string {
	return v.text
}

func (v *textBox) AddChild(childView View) {
	panic("not allowed")
}

func (v *textBox) SetBackgroundColor(c Color) {
	v.backgroundColor = c
}

func (v *textBox) BackgroundColor() Color {
	return v.backgroundColor
}

func (v *textBox) Height() int {
	return 1
}

func (v *textBox) Width() int {
	if v.width == -1 {
		return v.containerWidth
	}

	return v.width
}

func (v *textBox) SetWidth(width int) {
	v.width = width
}

func (v *textBox) setContainerWidth(width int) {
	v.containerWidth = width
}

func (v *textBox) Render() [][]Pixel {
	rows := NewPixels(v.Height(), v.Width(), v.backgroundColor)

	for i, c := range v.text {
		if i >= v.Width() {
			// The text exceeds the view, it will be hidden.
			break
		}

		rows[0][i].Character = c
	}

	return rows
}
