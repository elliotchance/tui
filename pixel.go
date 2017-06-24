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
				Character:       0,
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
	case bColor == NoColor && tColor == NoColor && bText == 0 && tText == 0,
		bColor == NoColor && tColor != NoColor && bText == 0 && tText == 0,
		bColor == NoColor && tColor != NoColor && bText == 0 && tText != 0,
		bColor == NoColor && tColor != NoColor && bText != 0 && tText == 0,
		bColor == NoColor && tColor != NoColor && bText != 0 && tText != 0:
		return Pixel{
			BackgroundColor: tColor,
			Character:       tText,
		}
	case bColor == NoColor && tColor == NoColor && bText == 0 && tText != 0,
		bColor == NoColor && tColor == NoColor && bText != 0 && tText != 0,
		bColor != NoColor && tColor == NoColor && bText == 0 && tText != 0,
		bColor != NoColor && tColor == NoColor && bText != 0 && tText != 0:
		return Pixel{
			BackgroundColor: bColor,
			Character:       tText,
		}
	case bColor != NoColor && tColor != NoColor && bText == 0 && tText == 0,
		bColor != NoColor && tColor != NoColor && bText == 0 && tText != 0,
		bColor != NoColor && tColor != NoColor && bText != 0 && tText != 0,
		bColor != NoColor && tColor != NoColor && bText != 0 && tText == 0:
		return Pixel{
			BackgroundColor: tColor,
			Character:       tText,
		}
	case bColor == NoColor && tColor == NoColor && bText != 0 && tText == 0,
		bColor != NoColor && tColor == NoColor && bText == 0 && tText == 0,
		bColor != NoColor && tColor == NoColor && bText != 0 && tText == 0:
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

func movePixelsRight(pixels [][]Pixel, rightAmount int) [][]Pixel {
	// FIXME: We assume that the shift right will be positive.

	height := len(pixels)
	width := len(pixels[0])
	rows := [][]Pixel{}

	for rowIndex := 0; rowIndex < height; rowIndex++ {
		// FIXME: This is a hack to make sure we always have room for the
		// rightAmount. It should pick the smallest value.
		row := make([]Pixel, width+rightAmount)
		for colIndex := 0; colIndex < rightAmount; colIndex++ {
			row[colIndex] = Pixel{
				Character:       0,
				BackgroundColor: NoColor,
			}
		}

		for colIndex := 0; colIndex < width; colIndex++ {
			row[colIndex+rightAmount] = pixels[rowIndex][colIndex]
		}

		rows = append(rows, row)
	}

	return rows
}

func movePixelsDown(pixels [][]Pixel, downAmount int) [][]Pixel {
	// FIXME: We assume that the shift down will be positive.

	height := len(pixels)
	width := len(pixels[0])
	rows := [][]Pixel{}

	// Create blank rows for the shift down.
	for rowIndex := 0; rowIndex < downAmount; rowIndex++ {
		row := make([]Pixel, width)
		for colIndex := 0; colIndex < width; colIndex++ {
			row[colIndex].Character = 0
		}

		rows = append(rows, row)
	}

	for rowIndex := 0; rowIndex < height; rowIndex++ {
		row := make([]Pixel, width)
		for colIndex := 0; colIndex < width; colIndex++ {
			row[colIndex] = pixels[rowIndex][colIndex]
		}

		rows = append(rows, row)
	}

	return rows
}
