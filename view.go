package tui

type Renderer interface {
	Render() [][]Pixel
}

type HasBackgroundColor interface {
	SetBackgroundColor(Color)
	BackgroundColor() Color
}

type View interface {
	Renderer
	HasBackgroundColor

	Size() MutableSize

	// Subviews
	AddTextBox(string) TextBox
}

type view struct {
	backgroundColor Color
	child           TextBox
	size            MutableSize
}

func (v *view) SetBackgroundColor(c Color) {
	v.backgroundColor = c
}

func (v *view) BackgroundColor() Color {
	return v.backgroundColor
}

func (v *view) Render() [][]Pixel {
	size := v.Size()
	renderedView := NewPixels(size.Height(), size.Width(), v.backgroundColor)

	if v.child != nil {
		renderedChild := v.child.Render()
		renderedView = OverlayPixels(renderedView, renderedChild)
	}

	return renderedView
}

func (v *view) Size() MutableSize {
	return v.size
}
