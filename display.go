package tui

import (
	"fmt"
)

func Display(pixels [][]Pixel) {
	for row := 0; row < len(pixels); row++ {
		for col := 0; col < len(pixels[0]); col++ {
			pixel := pixels[row][col]

			if pixel.BackgroundColor != NoColor {
				fmt.Printf("\x1b[48;5;%dm%c\x1b[0m", pixel.BackgroundColor, pixel.Character)
			} else {
				fmt.Printf("%c", pixel.Character)
			}
		}
		fmt.Printf("\n")
	}
}
