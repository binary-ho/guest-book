package rgb

type RGB struct {
	Red, Green, Blue value
}

// TODO : 일단은 너무 예민하게 용량을 줄이려 하지 말자
type value uint8
type Color uint32

func (rgb *RGB) GetHexColor() Color {
	red := uint32(rgb.Red) << 16
	green := uint32(rgb.Green) << 8
	blue := uint32(rgb.Blue)
	return Color(red | green | blue)
}
