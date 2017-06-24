package tui

type Sizer interface {
	Height() int
	Width() int
	HeightIsFlexible() bool
	WidthIsFlexible() bool
	HeightPercentage() float64
	WidthPercentage() float64
	AbsoluteLeft() int
	AbsoluteTop() int
	setContainerSize(int, int, int, int)
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
	containerLeft, containerTop       int
}

func newMutableSize(left, top, containerHeight, containerWidth int) *Size {
	return &Size{
		height:           1.0,
		width:            1.0,
		heightIsFlexible: true,
		widthIsFlexible:  true,
		containerHeight:  containerHeight,
		containerWidth:   containerWidth,
		containerLeft:    left,
		containerTop:     top,
	}
}

func (s *Size) AbsoluteLeft() int {
	return s.containerLeft
}

func (s *Size) AbsoluteTop() int {
	return s.containerTop
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

func (s *Size) HeightPercentage() float64 {
	if s.heightIsFlexible {
		return s.height
	}

	return s.height / float64(s.containerHeight)
}

func (s *Size) WidthPercentage() float64 {
	if s.widthIsFlexible {
		return s.width
	}

	return s.width / float64(s.containerWidth)
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

func (s *Size) setContainerSize(left, top, height, width int) {
	s.containerLeft = left
	s.containerTop = top
	s.containerHeight = height
	s.containerWidth = width
}
