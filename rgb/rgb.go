package rgb

type RGB struct {
	Red, Green, Blue value
}

type value uint8
type Color uint32

func (rgb *RGB) GetHexColor() Color {
	red := uint32(rgb.Red) << 16
	green := uint32(rgb.Green) << 8
	blue := uint32(rgb.Blue)
	return Color(red | green | blue)
}
