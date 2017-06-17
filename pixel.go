package tui

type Pixel struct {
	Character       rune
	BackgroundColor Color
}

func NewPixels(height, width int, backgroundColor Color) [][]Pixel {
	rows := [][]Pixel{}

	for row := 0; row < height; row++ {
		row := make([]Pixel, width)
		for col := 0; col < width; col++ {
			row[col] = Pixel{
				Character:       ' ',
				BackgroundColor: backgroundColor,
			}
		}

		rows = append(rows, row)
	}

	return rows
}

func OverlayPixel(bottom, top Pixel) (p Pixel) {
	bColor, bText := bottom.BackgroundColor, bottom.Character
	tColor, tText := top.BackgroundColor, top.Character

	switch {
	case bColor == NoColor && tColor == NoColor && bText == ' ' && tText == ' ',
		bColor == NoColor && tColor != NoColor && bText == ' ' && tText == ' ',
		bColor == NoColor && tColor != NoColor && bText == ' ' && tText != ' ',
		bColor == NoColor && tColor != NoColor && bText != ' ' && tText == ' ',
		bColor == NoColor && tColor != NoColor && bText != ' ' && tText != ' ':
		return Pixel{
			BackgroundColor: tColor,
			Character:       tText,
		}
	case bColor == NoColor && tColor == NoColor && bText == ' ' && tText != ' ',
		bColor == NoColor && tColor == NoColor && bText != ' ' && tText != ' ',
		bColor != NoColor && tColor == NoColor && bText == ' ' && tText != ' ',
		bColor != NoColor && tColor == NoColor && bText != ' ' && tText != ' ',
		bColor != NoColor && tColor != NoColor && bText == ' ' && tText == ' ',
		bColor != NoColor && tColor != NoColor && bText == ' ' && tText != ' ',
		bColor != NoColor && tColor != NoColor && bText != ' ' && tText != ' ',
		bColor != NoColor && tColor != NoColor && bText != ' ' && tText == ' ':
		return Pixel{
			BackgroundColor: bColor,
			Character:       tText,
		}
	case bColor == NoColor && tColor == NoColor && bText != ' ' && tText == ' ',
		bColor != NoColor && tColor == NoColor && bText == ' ' && tText == ' ',
		bColor != NoColor && tColor == NoColor && bText != ' ' && tText == ' ':
		return Pixel{
			BackgroundColor: bColor,
			Character:       bText,
		}
	}

	return
}

func OverlayPixels(bottom, top [][]Pixel) [][]Pixel {
	// FIXME: We will assume that top must be smaller in both dimensions than
	// bottom.
	height := len(bottom)
	width := len(bottom[0])

	rows := [][]Pixel{}

	for rowIndex := 0; rowIndex < height; rowIndex++ {
		row := make([]Pixel, width)
		for colIndex := 0; colIndex < width; colIndex++ {
			row[colIndex] = bottom[rowIndex][colIndex]
			if rowIndex < len(top) && colIndex < len(top[rowIndex]) {
				row[colIndex] = OverlayPixel(bottom[rowIndex][colIndex], top[rowIndex][colIndex])
			}
		}

		rows = append(rows, row)
	}

	return rows
}
