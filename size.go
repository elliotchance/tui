package tui

type Size interface {
	Height() int
	Width() int
	HeightIsFlexible() bool
	WidthIsFlexible() bool
}

type MutableSize interface {
	Size
	SetHeight(int)
	SetWidth(int)
	SetFlexibleHeight(float64)
	SetFlexibleWidth(float64)
}

type size struct {
	height, width                     float64
	heightIsFlexible, widthIsFlexible bool
	containerHeight, containerWidth   int
}

func newMutableSize(containerHeight, containerWidth int) MutableSize {
	return &size{
		height:           1.0,
		width:            1.0,
		heightIsFlexible: true,
		widthIsFlexible:  true,
		containerHeight:  containerHeight,
		containerWidth:   containerWidth,
	}
}

func (s *size) Height() int {
	if s.heightIsFlexible {
		return int(s.height * float64(s.containerHeight))
	}

	return int(s.height)
}

func (s *size) Width() int {
	if s.widthIsFlexible {
		return int(s.width * float64(s.containerWidth))
	}

	return int(s.width)
}

func (s *size) SetHeight(height int) {
	s.heightIsFlexible = false
	s.height = float64(height)
}

func (s *size) SetWidth(width int) {
	s.widthIsFlexible = false
	s.width = float64(width)
}

func (s *size) SetFlexibleHeight(flexibleHeight float64) {
	s.heightIsFlexible = true
	s.height = flexibleHeight
}

func (s *size) SetFlexibleWidth(flexibleWidth float64) {
	s.widthIsFlexible = true
	s.width = flexibleWidth
}

func (s *size) HeightIsFlexible() bool {
	return s.heightIsFlexible
}

func (s *size) WidthIsFlexible() bool {
	return s.widthIsFlexible
}
