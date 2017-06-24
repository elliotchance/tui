package tui

type BoxView struct {
	backgroundColor Color
	size            *Size
	view            *View
	title           string
}

func (v *View) AddBoxView() *BoxView {
	boxView := &BoxView{
		backgroundColor: NoColor,
		size:            newMutableSize(v.Size().Height(), v.Size().Width()),
		view:            newView(v.Size().Height()-2, v.Size().Width()-2),
		title:           "",
	}

	v.child = boxView

	return boxView
}

func (v *BoxView) SetTitle(title string) {
	v.title = title
}

func (v *BoxView) Title() string {
	return v.title
}

func (v *BoxView) View() *View {
	return v.view
}

func (v *BoxView) SetBackgroundColor(c Color) {
	v.backgroundColor = c
}

func (v *BoxView) BackgroundColor() Color {
	return v.backgroundColor
}

func (v *BoxView) Size() MutableSizer {
	return v.size
}

func (v *BoxView) Render() [][]Pixel {
	width := v.Size().Width()
	height := v.Size().Height()
	rows := NewPixels(height, width, v.backgroundColor)

	// Draw the title
	rows[0][0].Character = '┌'
	if v.title == "" {
		for i := 1; i < width-1; i++ {
			rows[0][i].Character = '─'
		}
	} else {
		rows[0][1].Character = ' '
		for i := 0; i < len(v.title); i++ {
			rows[0][i+2].Character = rune(v.title[i])
		}
		rows[0][len(v.title)+2].Character = ' '
		for i := len(v.title) + 3; i < width-1; i++ {
			rows[0][i].Character = '─'
		}
	}
	rows[0][width-1].Character = '┐'

	// Draw the rest of the box
	for i := 1; i < height-1; i++ {
		rows[i][0].Character = '│'
		rows[i][width-1].Character = '│'
	}
	rows[height-1][0].Character = '└'
	rows[height-1][width-1].Character = '┘'
	for i := 1; i < width-1; i++ {
		rows[height-1][i].Character = '─'
	}

	viewRendered := v.View().Render()
	viewRendered = movePixelsRight(viewRendered, 1)
	viewRendered = movePixelsDown(viewRendered, 1)

	rows = OverlayPixels(rows, viewRendered)

	return rows
}

func (v *BoxView) setContainerSize(height, width int) {
	v.Size().setContainerSize(height, width)
	v.view.setContainerSize(height, width)
}
