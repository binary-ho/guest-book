package rgb

type RGB struct {
	Red, Green, Blue uint
}

func (rgb *RGB) Get() uint {
	return rgb.Red
}

type Getter interface {
	Get() uint
}
