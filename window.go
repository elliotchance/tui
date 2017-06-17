package tui

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Pixel byte

type Renderer interface {
	Render() [][]Pixel
}

type Window interface {
	Renderer

	Height() int
	Width() int
}

type window struct {
	height, width int
}

func getTerminalSize() (height, width int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Sscanf(string(out), "%d %d", &height, &width)
	return
}

func (w *window) Height() int {
	return w.height
}

func (w *window) Width() int {
	return w.width
}

func NewPixels(height, width int) [][]Pixel {
	rows := [][]Pixel{}

	for row := 0; row < height; row++ {
		row := make([]Pixel, width)
		for col := 0; col < width; col++ {
			row[col] = ' '
		}

		rows = append(rows, row)
	}

	return rows
}

func (w *window) Render() [][]Pixel {
	rows := NewPixels(w.height, w.width)

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
	height, width := getTerminalSize()
	return &window{height, width}
}
