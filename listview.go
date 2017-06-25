package tui

import (
	"github.com/nsf/termbox-go"
)

type ListView struct {
	backgroundColor Color
	size            *Size
	items           []string
	selectedIndex   int

	// Events
	onIndexChanged func()
}

func (v *View) AddListView() *ListView {
	listView := &ListView{
		backgroundColor: NoColor,
		size:            newMutableSize(v.Size().AbsoluteLeft(), v.Size().AbsoluteTop(), v.Size().Height(), v.Size().Width()),
		items:           []string{},
		selectedIndex:   -1,
	}

	v.child = listView

	return listView
}

func (v *ListView) Items() []string {
	return v.items
}

func (v *ListView) SetItems(items []string) {
	v.items = items

	if v.selectedIndex < 0 && len(items) > 0 {
		// This happens when we are receiving items when there were previously
		// none. This will cause the event OnIndexChanged to fire.
		v.SetSelectedIndex(0)
	}
}

func (v *ListView) SetBackgroundColor(c Color) {
	v.backgroundColor = c
}

func (v *ListView) BackgroundColor() Color {
	return v.backgroundColor
}

func (v *ListView) Size() MutableSizer {
	return v.size
}

func (v *ListView) Render() [][]Pixel {
	height := v.Size().Height()
	width := v.Size().Width()

	rows := NewPixels(height, width, v.backgroundColor)

	for line := 0; line < len(v.items); line++ {
		if line >= height {
			break
		}

		for c := 0; c < len(v.items[line]); c++ {
			if c >= width {
				break
			}

			rows[line][c].Character = rune(v.items[line][c])
		}
	}

	// Highlight the selected item.
	if v.selectedIndex >= 0 && v.selectedIndex < len(rows) {
		for c := 0; c < width; c++ {
			rows[v.selectedIndex][c].BackgroundColor = Blue
		}
	}

	return rows
}

func (v *ListView) SelectedIndex() int {
	return v.selectedIndex
}

func (v *ListView) SetSelectedIndex(selectedIndex int) {
	// Make sure they do not go beyond the bounds of the list.
	if selectedIndex < 0 {
		selectedIndex = 0
	}
	if selectedIndex >= len(v.items) {
		selectedIndex = len(v.items) - 1
	}

	shouldFireEvent := v.selectedIndex != selectedIndex
	v.selectedIndex = selectedIndex

	if shouldFireEvent && v.onIndexChanged != nil {
		v.onIndexChanged()
	}
}

func (v *ListView) setContainerSize(left, top, height, width int) {
	v.Size().setContainerSize(left, top, height, width)
}

func (v *ListView) getViewForPosition(x, y int) Renderer {
	return v
}

func (v *ListView) handleEvent(e termbox.Event) {
	switch e.Type {
	case termbox.EventKey:
		switch e.Key {
		case termbox.KeyArrowDown:
			v.SetSelectedIndex(v.SelectedIndex() + 1)

		case termbox.KeyArrowUp:
			v.SetSelectedIndex(v.SelectedIndex() - 1)
		}

	case termbox.EventMouse:
		newSelectedIndex := e.MouseY - v.Size().AbsoluteTop()
		if newSelectedIndex < 0 || newSelectedIndex >= len(v.Items()) {
			break
		}

		v.SetSelectedIndex(newSelectedIndex)
	}
}

func (v *ListView) OnIndexChanged(onIndexChanged func()) {
	v.onIndexChanged = onIndexChanged
}
