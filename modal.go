package tui

type Modal struct {
	backgroundColor Color
	size            MutableSizer
	view            *View
}

func NewModal(height, width int) *Modal {
	return &Modal{
		backgroundColor: Cyan,
		size:            newMutableSize(0, 0, height, width),
		view:            newView(0, 0, height, width),
	}
}

func (m *Modal) View() *View {
	return m.view
}

func (m *Modal) Size() Sizer {
	return m.size
}

func (m *Modal) Render() [][]Pixel {
	renderedView := m.View().Render()
	size := m.Size()
	renderedWindow := NewPixels(size.Height(), size.Width(), m.backgroundColor)

	return OverlayPixels(renderedWindow, renderedView)
}

func (m *Modal) setContainerSize(left, top, height, width int) {
	m.Size().setContainerSize(left, top, height, width)
	m.View().setContainerSize(left, top, height, width)
}
