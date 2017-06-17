package tui

type Renderer interface {
	Render() [][]Pixel
}

type View interface {
	Renderer

	// Size
	Height() int
	Width() int

	// Colors
	SetBackgroundColor(Color)
	BackgroundColor() Color
}

type view struct {
	height, width   int
	backgroundColor Color
}

func (v *view) Height() int {
	return v.height
}

func (v *view) Width() int {
	return v.width
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
