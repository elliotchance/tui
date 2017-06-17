package tui

type TextBox interface {
	View

	SetText(string)
	Text() string
}

type textBox struct {
	text            string
	backgroundColor Color
}

func NewTextBox(text string) TextBox {
	return &textBox{
		text:            text,
		backgroundColor: NoColor,
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
	return 20
}

func (v *textBox) Render() [][]Pixel {
	rows := NewPixels(v.Height(), v.Width(), v.backgroundColor)

	for i, c := range v.text {
		rows[0][i].Character = c
	}

	return rows
}
