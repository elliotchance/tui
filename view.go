package tui

import "github.com/nsf/termbox-go"

type Renderer interface {
	Render() [][]Pixel
	setContainerSize(left, top, height, width int)
	getViewForPosition(x, y int) Renderer
	handleEvent(e termbox.Event)
}

type View struct {
	backgroundColor Color
	child           Renderer
	size            MutableSizer
}

func newView(left, top, height, width int) *View {
	return &View{
		size:            newMutableSize(left, top, height, width),
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

func (v *View) setContainerSize(left, top, height, width int) {
	v.Size().setContainerSize(left, top, height, width)
	if v.child != nil {
		v.child.setContainerSize(left, top, height, width)
	}
}

func (v *View) getViewForPosition(x, y int) Renderer {
	if v.child != nil {
		return v.child.getViewForPosition(x, y)
	}

	return v
}

func (v *View) handleEvent(e termbox.Event) {
}
