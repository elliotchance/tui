package tui

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
		size:            newMutableSize(1 /*v.Size().Height()*/, v.Size().Width()),
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

func (v *textBox) Render() [][]Pixel {
	rows := NewPixels(v.Size().Height(), v.Size().Width(), v.backgroundColor)

	for i, c := range v.text {
		if i >= v.Size().Width() {
			// The text exceeds the view, it will be hidden.
			break
		}

		rows[0][i].Character = c
	}

	return rows
}
