package tui

type ListView struct {
	backgroundColor Color
	size            *Size
	items           []string
}

func (v *View) AddListView() *ListView {
	listView := &ListView{
		backgroundColor: NoColor,
		size:            newMutableSize(v.Size().Height(), v.Size().Width()),
		items:           []string{},
	}

	v.child = listView

	return listView
}

func (v *ListView) Items() []string {
	return v.items
}

func (v *ListView) SetItems(items []string) {
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

	return rows
}
