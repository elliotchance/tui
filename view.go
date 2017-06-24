package tui

type Renderer interface {
	Render() [][]Pixel
	setContainerSize(int, int)
}

type View struct {
	backgroundColor Color
	child           Renderer
	size            MutableSizer
}

func newView(height, width int) *View {
	return &View{
		size:            newMutableSize(height, width),
		backgroundColor: NoColor,
	}
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

func (v *View) setContainerSize(height, width int) {
	v.Size().setContainerSize(height, width)
	if v.child != nil {
		v.child.setContainerSize(height, width)
	}
}
