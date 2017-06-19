package tui

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Window struct {
	view            *View
	backgroundColor Color
	size            *Size
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

func (w *Window) View() *View {
	return w.view
}

func (w *Window) Size() Sizer {
	return w.size
}

func (w *Window) Render() [][]Pixel {
	renderedView := w.View().Render()

	size := w.Size()
	renderedWindow := NewPixels(size.Height(), size.Width(), w.backgroundColor)

	return OverlayPixels(renderedWindow, renderedView)
}

func (w *Window) SetBackgroundColor(c Color) {
	w.backgroundColor = c
}

func (w *Window) BackgroundColor() Color {
	return w.backgroundColor
}

func newWindow(height, width int) *Window {
	return &Window{
		size:            newMutableSize(height, width),
		backgroundColor: NoColor,
		view:            newView(height, width),
	}
}

func MainWindow() *Window {
	height, width := getTerminalSize()

	return newWindow(height, width)
}
