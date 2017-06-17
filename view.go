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

	AddChild(view View)
}

type ResizableView interface {
	View

	SetHeight(int)
	SetWidth(int)
	SetFlexibleHeight(float64)
	SetFlexibleWidth(float64)
}

type hasContainerWidth interface {
	setContainerWidth(int)
}

type view struct {
	height, width                     float64
	heightIsFlexible, widthIsFlexible bool
	backgroundColor                   Color
	containerHeight, containerWidth   int
	child                             View
}

func (v *view) Height() int {
	if v.heightIsFlexible {
		return int(v.height * float64(v.containerHeight))
	}

	return int(v.height)
}

func (v *view) Width() int {
	if v.widthIsFlexible {
		return int(v.width * float64(v.containerWidth))
	}

	return int(v.width)
}

func (v *view) SetHeight(height int) {
	v.heightIsFlexible = false
	v.height = float64(height)
}

func (v *view) SetWidth(width int) {
	v.widthIsFlexible = false
	v.width = float64(width)
}

func (v *view) SetFlexibleHeight(flexibleHeight float64) {
	v.heightIsFlexible = true
	v.height = flexibleHeight
}

func (v *view) SetFlexibleWidth(flexibleWidth float64) {
	v.widthIsFlexible = true
	v.width = flexibleWidth
}

func (v *view) SetBackgroundColor(c Color) {
	v.backgroundColor = c
}

func (v *view) BackgroundColor() Color {
	return v.backgroundColor
}

func (v *view) AddChild(childView View) {
	if c, ok := childView.(hasContainerWidth); ok {
		c.setContainerWidth(v.Width())
	}

	v.child = childView
}

func (v *view) Render() [][]Pixel {
	renderedView := NewPixels(v.Height(), v.Width(), v.backgroundColor)

	if v.child != nil {
		renderedChild := v.child.Render()
		renderedView = OverlayPixels(renderedView, renderedChild)
	}

	return renderedView
}
