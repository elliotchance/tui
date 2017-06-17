package tui

// Color represents an 8-bit color with 3 bits for red and green and 2 bits for
// blue.
type Color uint8

func round(x float64) int {
	return int(x + 0.5)
}

func NewColorRGB(r, g, b float64) Color {
	// This code has been adpated from:
	// https://gist.github.com/MicahElliott/719710#gistcomment-615676
	return Color(round(36*(r*5) + 6*(g*5) + (b * 5) + 16))
}
