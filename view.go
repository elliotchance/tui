package tui

type Renderer interface {
	Render() [][]Pixel
}

type HasBackgroundColor interface {
	SetBackgroundColor(Color)
	BackgroundColor() Color
}

type HasSize interface {
	Height() int
	Width() int
}

type View interface {
	Renderer
	HasBackgroundColor
	HasSize

	SetHeight(int)
	SetWidth(int)
}

type view struct {
	height, width   int
	backgroundColor Color
}

func (v *view) Height() int {
	return v.height
}

func (v *view) SetHeight(height int) {
	v.height = height
}

func (v *view) Width() int {
	return v.width
}

func (v *view) SetWidth(width int) {
	v.width = width
}

func (v *view) SetBackgroundColor(c Color) {
	v.backgroundColor = c
}

func (v *view) BackgroundColor() Color {
	return v.backgroundColor
}

func (v *view) Render() [][]Pixel {
	rows := NewPixels(v.height, v.width, v.backgroundColor)

	return rows
}
