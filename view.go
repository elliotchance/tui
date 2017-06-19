package tui

type Renderer interface {
	Render() [][]Pixel
}

type View struct {
	backgroundColor Color
	child           Renderer
	size            *Size
}

func (v *View) SetBackgroundColor(c Color) {
	v.backgroundColor = c
}

func (v *View) BackgroundColor() Color {
	return v.backgroundColor
}

func (v *View) Render() [][]Pixel {
	size := v.Size()
	renderedView := NewPixels(size.Height(), size.Width(), v.backgroundColor)

	if v.child != nil {
		renderedChild := v.child.Render()
		renderedView = OverlayPixels(renderedView, renderedChild)
	}

	return renderedView
}

func (v *View) Size() MutableSizer {
	return v.size
}
