package tui

import "github.com/nsf/termbox-go"

type SplitView struct {
	backgroundColor     Color
	size                *Size
	leftView, rightView *View
}

func (v *View) AddSplitView(leftViewWidth int) *SplitView {
	viewWidth := v.Size().Width()
	rightViewWidth := viewWidth - leftViewWidth

	splitView := &SplitView{
		backgroundColor: NoColor,
		size:            newMutableSize(v.Size().AbsoluteLeft(), v.Size().AbsoluteTop(), v.Size().Height(), viewWidth),
		leftView:        newView(v.Size().AbsoluteLeft(), v.Size().AbsoluteTop(), v.Size().Height(), leftViewWidth),
		rightView:       newView(v.Size().AbsoluteLeft()+leftViewWidth, v.Size().AbsoluteTop(), v.Size().Height(), rightViewWidth),
	}

	v.child = splitView

	return splitView
}

func (v *View) AddFlexibleSplitView(leftWidthPercentage float64) *SplitView {
	viewWidth := v.Size().Width()
	leftViewWidth := int(leftWidthPercentage * float64(viewWidth))
	rightViewWidth := viewWidth - leftViewWidth

	splitView := &SplitView{
		backgroundColor: NoColor,
		size:            newMutableSize(v.Size().AbsoluteLeft(), v.Size().AbsoluteTop(), v.Size().Height(), viewWidth),
		leftView:        newView(v.Size().AbsoluteLeft(), v.Size().AbsoluteTop(), v.Size().Height(), leftViewWidth),
		rightView:       newView(v.Size().AbsoluteLeft()+leftViewWidth, v.Size().AbsoluteTop(), v.Size().Height(), rightViewWidth),
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
	leftWidth := v.LeftView().Size().Width()
	renderedRight := movePixelsRight(v.RightView().Render(), leftWidth)

	rows = OverlayPixels(rows, renderedLeft)
	rows = OverlayPixels(rows, renderedRight)

	return rows
}

func (v *SplitView) setContainerSize(left, top, height, width int) {
	v.Size().setContainerSize(left, top, height, width)

	leftView := v.LeftView()
	rightView := v.RightView()

	// TODO: This only works for a fixed width splitview.
	leftView.setContainerSize(left, top, height, leftView.Size().Width())
	rightView.setContainerSize(left+leftView.Size().Width(), top, height, width-leftView.Size().Width())
}

func (v *SplitView) getViewForPosition(x, y int) Renderer {
	if x >= v.RightView().Size().AbsoluteLeft() {
		return v.RightView().getViewForPosition(x, y)
	}

	return v.LeftView().getViewForPosition(x, y)
}

func (v *SplitView) handleEvent(e termbox.Event) {
}
