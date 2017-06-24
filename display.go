package tui

import (
	"fmt"
)

func Display(pixels [][]Pixel) string {
	out := ""
	for row := 0; row < len(pixels); row++ {
		for col := 0; col < len(pixels[0]); col++ {
			pixel := pixels[row][col]

			if pixel.Character == 0 {
				pixel.Character = ' '
			}

			if pixel.BackgroundColor != NoColor {
				out += fmt.Sprintf("\x1b[48;5;%dm%c\x1b[0m", pixel.BackgroundColor, pixel.Character)
			} else {
				out += fmt.Sprintf("%c", pixel.Character)
			}
		}
		out += "\n"
	}

	return out
}

func Print(pixels [][]Pixel) {
	fmt.Print(Display(pixels))
}
