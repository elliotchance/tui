package tui

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"log"
	"os"
	"os/exec"
)

type Window struct {
	view            *View
	backgroundColor Color
	size            *Size
	activeColor     Color
	activeView      Renderer
	modal           *Modal
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
	// Render the window first
	renderedView := w.View().Render()
	size := w.Size()

	renderedWindow := NewPixels(size.Height(), size.Width(), w.backgroundColor)
	renderedWindow = OverlayPixels(renderedWindow, renderedView)

	// Now render the modal over the top
	if w.modal != nil {
		renderedModal := w.modal.Render()

		// Move the modal so it appears at the same coordinates as it stores
		// internally.
		renderedModal = movePixelsRight(renderedModal, w.modal.Size().AbsoluteLeft())
		renderedModal = movePixelsDown(renderedModal, w.modal.Size().AbsoluteTop())

		renderedWindow = OverlayPixels(renderedWindow, renderedModal)
	}

	return renderedWindow
}

func (w *Window) SetBackgroundColor(c Color) {
	w.backgroundColor = c
}

func (w *Window) BackgroundColor() Color {
	return w.backgroundColor
}

func (w *Window) SetActiveView(activeView Renderer) {
	w.activeView = activeView
}

func (w *Window) ActiveView() Renderer {
	return w.activeView
}

func (w *Window) Start() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.SetOutputMode(termbox.Output256)
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
				goto exit
			}

			w.activeView.handleEvent(ev)

		case termbox.EventMouse:
			activeView := w.View().getViewForPosition(ev.MouseX, ev.MouseY)
			w.SetActiveView(activeView)
			activeView.handleEvent(ev)

		case termbox.EventResize:
			// Set the new dimensions of the main window.
			height, width := getTerminalSize()
			w.size.height = float64(height)
			w.size.width = float64(width)

			// Now we have to cause all subviews to resize.
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			w.view.setContainerSize(0, 0, height, width)
		}
	}

exit:
}

func newWindow(height, width int) *Window {
	view := newView(0, 0, height, width)

	return &Window{
		size:            newMutableSize(0, 0, height, width),
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

func (w *Window) ShowModal(modal *Modal) {
	// When the modal is added we need to adjust the position so that it appears
	// in the center of the screen.
	left := (w.Size().Width() - modal.Size().Width()) / 2
	top := (w.Size().Height() - modal.Size().Height()) / 2

	modal.setContainerSize(left, top, modal.Size().Height(), modal.Size().Width())

	w.modal = modal
}
