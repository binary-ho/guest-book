package canvas

import (
	"guestbook/common/coord"
	"guestbook/entity"
	"guestbook/entity/user"
	"guestbook/register"
)

type Entity struct {
	id     entity.ID
	owner  user.Entity
	size   size
	colors register.Colors
}

type size struct {
	height coord.Y
	width  coord.X
}

func (canvas *Entity) Merge(y coord.Y, x coord.X, state *register.State) *register.LWWRegister {
	key := canvas.getKey(y, x)
	color, exist := canvas.colors[key]

	if !exist {
		canvas.colors[key] = *register.CreateLLWWRegister(state)
		color := canvas.colors[key]
		return &color
	}
	return color.Merge(state)
}

func (canvas *Entity) getKey(y coord.Y, x coord.X) register.ColorKey {
	width := canvas.size.width
	height := canvas.size.height
	// TODO : 예외 발생으로 변경
	if y >= height || x >= width {
		return 0
	}

	position := calculatePosition(width, y, x)
	return register.ColorKey(position)
}

func calculatePosition(width coord.X, y coord.Y, x coord.X) uint32 {
	a := uint32(width) * uint32(y)
	b := uint32(x)
	return a + b
}
