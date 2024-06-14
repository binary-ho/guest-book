package register

import (
	"guestbook/common/coord"
	"guestbook/entity"
)

type Canvas struct {
	id   entity.ID
	size size
	data colors
}

type size struct {
	height coord.Y
	width  coord.X
}

func (canvas *Canvas) Merge(y coord.Y, x coord.X, state *State) {
	key := canvas.getKey(y, x)
	register, exist := canvas.data[key]

	if !exist {
		canvas.data[key] = LWWRegister{state: state}
		return
	}
	register.Merge(state)
}

func (canvas *Canvas) getKey(y coord.Y, x coord.X) colorKey {
	width := canvas.size.width
	height := canvas.size.height
	// TODO : 예외 발생으로 변경
	if y >= height || x >= width {
		return 0
	}

	position := calculatePosition(width, y, x)
	return colorKey(position)
}

func calculatePosition(width coord.X, y coord.Y, x coord.X) uint32 {
	a := uint32(width) * uint32(y)
	b := uint32(x)
	return a + b
}
