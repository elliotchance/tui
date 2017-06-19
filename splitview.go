package tui

type SplitView struct {
	backgroundColor     Color
	size                *Size
	leftView, rightView *View
}

func (v *View) AddSplitView() *SplitView {
	viewWidth := v.Size().Width()
	leftViewWidth := int(viewWidth / 2)
	rightViewWidth := viewWidth - leftViewWidth

	splitView := &SplitView{
		backgroundColor: NoColor,
		size:            newMutableSize(v.Size().Height(), viewWidth),
		leftView:        newView(v.Size().Height(), leftViewWidth),
		rightView:       newView(v.Size().Height(), rightViewWidth),
	}

	v.child = splitView

	return splitView
}

func (v *SplitView) LeftView() *View {
	return v.leftView
}

func (v *SplitView) RightView() *View {
	return v.rightView
}

func (v *SplitView) SetBackgroundColor(c Color) {
	v.backgroundColor = c
}

func (v *SplitView) BackgroundColor() Color {
	return v.backgroundColor
}

func (v *SplitView) Size() MutableSizer {
	return v.size
}

func (v *SplitView) Render() [][]Pixel {
	rows := NewPixels(v.Size().Height(), v.Size().Width(), v.backgroundColor)

	renderedLeft := v.LeftView().Render()
	renderedRight := movePixelsRight(v.RightView().Render(), 2)

	rows = OverlayPixels(rows, renderedLeft)
	rows = OverlayPixels(rows, renderedRight)

	return rows
}
