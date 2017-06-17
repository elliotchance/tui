package tui

// Color represents an 8-bit color with 3 bits for red and green and 2 bits for
// blue.
type Color int16

const (
	// NoColor is a special color that will not apply any color. The color of
	// the terminal window will be used.
	NoColor = Color(-1)

	Black       = Color(0)
	DarkRed     = Color(1)
	DarkGreen   = Color(2)
	DarkYellow  = Color(3)
	DarkBlue    = Color(4)
	DarkMagenta = Color(5)
	DarkCyan    = Color(6)
	Gray        = Color(7)
	DarkGray    = Color(8)
	Red         = Color(9)
	Green       = Color(10)
	Yellow      = Color(11)
	Blue        = Color(12)
	Magenta     = Color(13)
	Cyan        = Color(14)
	White       = Color(15)
)

func round(x float64) int {
	return int(x + 0.5)
}

func NewColorRGB(r, g, b float64) Color {
	// This code has been adpated from:
	// https://gist.github.com/MicahElliott/719710#gistcomment-615676
	return Color(round(36*(r*5) + 6*(g*5) + (b * 5) + 16))
}
