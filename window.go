package tui

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Window interface {
	Renderer
	HasBackgroundColor

	View() View
	Size() Size
}

type window struct {
	view            View
	backgroundColor Color
	size            Size
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

func (w *window) View() View {
	return w.view
}

func (w *window) Size() Size {
	return w.size
}

func (w *window) Render() [][]Pixel {
	renderedView := w.View().Render()

	size := w.Size()
	renderedWindow := NewPixels(size.Height(), size.Width(), w.backgroundColor)

	return OverlayPixels(renderedWindow, renderedView)
}

func (w *window) SetBackgroundColor(c Color) {
	w.backgroundColor = c
}

func (w *window) BackgroundColor() Color {
	return w.backgroundColor
}

func newWindow(height, width int) Window {
	return &window{
		size:            newMutableSize(height, width),
		backgroundColor: NoColor,

		view: &view{
			size:            newMutableSize(height, width),
			backgroundColor: NoColor,
		},
	}
}

func MainWindow() Window {
	height, width := getTerminalSize()

	return newWindow(height, width)
}
