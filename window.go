package tui

import "fmt"

type Pixel byte

type Renderer interface {
	Render() [][]Pixel
}

type Window interface {
	Renderer

	Width() int
	Height() int
}

type window struct {
	width, height int
}

func (w *window) Width() int {
	return w.width
}

func (w *window) Height() int {
	return w.height
}

func NewPixels(width, height int) [][]Pixel {
	rows := [][]Pixel{}

	for row := 0; row < height; row++ {
		row := make([]Pixel, width)
		for col := 0; col < width; col++ {
			row[col] = '.'
		}

		rows = append(rows, row)
	}

	return rows
}

func (w *window) Render() [][]Pixel {
	rows := NewPixels(w.width, w.height)

	return rows
}

func Display(pixels [][]Pixel) {
	for row := 0; row < len(pixels); row++ {
		for col := 0; col < len(pixels[0]); col++ {
			fmt.Printf("%c", pixels[row][col])
		}
		fmt.Printf("\n")
	}
}

func MainWindow() Window {
	return &window{5, 3}
}
