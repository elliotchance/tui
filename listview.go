package tui

import (
	"github.com/nsf/termbox-go"
)

type ListView struct {
	backgroundColor Color
	size            *Size
	items           []string
	selectedIndex   int
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
	if v.selectedIndex < 0 && len(items) > 0 {
		v.selectedIndex = 0
	}

	v.items = items
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
	v.selectedIndex = selectedIndex
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
		v.SetSelectedIndex(e.MouseY - v.Size().AbsoluteTop())
	}
}
