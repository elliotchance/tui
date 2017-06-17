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
	HasSize

	View() View
}

type window struct {
	view            View
	backgroundColor Color
	height, width   int
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

func (w *window) Render() [][]Pixel {
	view := w.View()
	renderedView := view.Render()

	// If the view has the same size as the window we do not need to do anything
	// else.
	if view.Height() == w.Height() && view.Width() == w.Width() {
		return renderedView
	}

	// Otherwise we have to render the window and overlay the view.
	renderedWindow := NewPixels(w.height, w.width, w.backgroundColor)

	return OverlayPixels(renderedWindow, renderedView)
}

func (w *window) SetBackgroundColor(c Color) {
	w.backgroundColor = c
}

func (w *window) BackgroundColor() Color {
	return w.backgroundColor
}

func (w *window) Height() int {
	return w.height
}

func (w *window) Width() int {
	return w.width
}

func MainWindow() Window {
	height, width := getTerminalSize()

	return &window{
		height:          height,
		width:           width,
		backgroundColor: NoColor,

		view: &view{
			height:           100.0,
			width:            100.0,
			heightIsFlexible: true,
			widthIsFlexible:  true,
			containerHeight:  height,
			containerWidth:   width,
			backgroundColor:  NoColor,
		},
	}
}
