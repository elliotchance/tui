package tui

type Pixel struct {
	Character       byte
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
