package tui

type Sizer interface {
	Height() int
	Width() int
	HeightIsFlexible() bool
	WidthIsFlexible() bool
}

type MutableSizer interface {
	Sizer
	SetHeight(int)
	SetWidth(int)
	SetFlexibleHeight(float64)
	SetFlexibleWidth(float64)
}

type Size struct {
	height, width                     float64
	heightIsFlexible, widthIsFlexible bool
	containerHeight, containerWidth   int
}

func newMutableSize(containerHeight, containerWidth int) *Size {
	return &Size{
		height:           1.0,
		width:            1.0,
		heightIsFlexible: true,
		widthIsFlexible:  true,
		containerHeight:  containerHeight,
		containerWidth:   containerWidth,
	}
}

func (s *Size) Height() int {
	if s.heightIsFlexible {
		return int(s.height * float64(s.containerHeight))
	}

	return int(s.height)
}

func (s *Size) Width() int {
	if s.widthIsFlexible {
		return int(s.width * float64(s.containerWidth))
	}

	return int(s.width)
}

func (s *Size) SetHeight(height int) {
	s.heightIsFlexible = false
	s.height = float64(height)
}

func (s *Size) SetWidth(width int) {
	s.widthIsFlexible = false
	s.width = float64(width)
}

func (s *Size) SetFlexibleHeight(flexibleHeight float64) {
	s.heightIsFlexible = true
	s.height = flexibleHeight
}

func (s *Size) SetFlexibleWidth(flexibleWidth float64) {
	s.widthIsFlexible = true
	s.width = flexibleWidth
}

func (s *Size) HeightIsFlexible() bool {
	return s.heightIsFlexible
}

func (s *Size) WidthIsFlexible() bool {
	return s.widthIsFlexible
}
