package tui

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Window interface {
	View() View
}

type window struct {
	view View
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

func MainWindow() Window {
	height, width := getTerminalSize()
	return &window{
		view: &view{
			height:          height,
			width:           width,
			backgroundColor: NoColor,
		},
	}
}
