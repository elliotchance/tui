package tui

import "github.com/nsf/termbox-go"

// Color represents an 8-bit color with 3 bits for red and green and 2 bits for
// blue.
type Color int16

const (
	// NoColor is a special color that will not apply any color. The color of
	// the parent view up to the terminal window itself will be used.
	NoColor = Color(termbox.ColorDefault)

	Black = Color(termbox.ColorBlack)
	//DarkRed     = Color(termbox.Attr)
	//DarkGreen   = Color(2)
	//DarkYellow  = Color(3)
	//DarkBlue    = Color(4)
	//DarkMagenta = Color(5)
	//DarkCyan    = Color(6)
	//Gray        = Color(termbox.Color)
	//DarkGray    = Color(8)
	Red     = Color(termbox.ColorRed)
	Green   = Color(termbox.ColorGreen)
	Yellow  = Color(termbox.ColorYellow)
	Blue    = Color(termbox.ColorBlue)
	Magenta = Color(termbox.ColorMagenta)
	Cyan    = Color(termbox.ColorCyan)
	White   = Color(termbox.ColorWhite)
)

func round(x float64) int {
	return int(x + 0.5)
}

func NewColorRGB(r, g, b float64) Color {
	// This code has been adpated from:
	// https://gist.github.com/MicahElliott/719710#gistcomment-615676
	return Color(round(36*(r*5) + 6*(g*5) + (b * 5) + 16))
}
