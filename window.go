package tui

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"github.com/nsf/termbox-go"
)

type Window struct {
	view            *View
	backgroundColor Color
	size            *Size
	activeColor     Color
	activeView      *View
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

func (w *Window) SetActiveView(activeView *View) {
	w.activeView = activeView
}

func (w *Window) ActiveView() *View {
	return w.activeView
}

func (w *Window) Start(h func(termbox.Event)) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for {
		pixels := w.Render()
		for rowId, row := range pixels {
			for colId, pixel := range row {
				termbox.SetCell(colId, rowId, pixel.Character, 0,
					termbox.Attribute(pixel.BackgroundColor))
			}
		}
		termbox.Flush()

		ev := termbox.PollEvent()

		switch ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlC {
				os.Exit(0)
			}

		case termbox.EventResize:
			// Set the new dimensions of the main window.
			height, width := getTerminalSize()
			w.size.height = float64(height)
			w.size.width = float64(width)

			// Now we have to cause all subviews to resize.
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			w.view.setContainerSize(height, width)
			continue
		}

		h(ev)
	}
}

func newWindow(height, width int) *Window {
	view := newView(height, width)

	return &Window{
		size:            newMutableSize(height, width),
		backgroundColor: NoColor,
		view:            view,
		activeColor:     Red,
		activeView:      view,
	}
}

func MainWindow() *Window {
	height, width := getTerminalSize()

	return newWindow(height, width)
}
