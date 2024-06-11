package register

import (
	"guestbook/entity"
)

type Canvas struct {
	id   entity.ID
	size size
	data registerMap
}

/*
TODO : (0, 0) 부터 시작할지 (1, 1) 부터 시작할지?
(0, 0)은 기본값이므로 문제가 생길 수도 있음.
일단은 (0, 0)으로 한다. (1, 1)의 경우 값을 클라이언트로 내려줄 때 까다로욱 수도 있음
*/
type registerMap map[uint32]LWWRegister

type size struct {
	width, height uint16
}

func (canvas *Canvas) Merge(y, x uint16, state *State) {
	key := canvas.getKey(y, x)
	register, exist := canvas.data[key]

	if !exist {
		canvas.data[key] = LWWRegister{state: state}
		return
	}
	register.Merge(state)
}

func (canvas *Canvas) getKey(y, x uint16) uint32 {
	width := canvas.size.width
	height := canvas.size.height
	// TODO : 예외 발생으로 변경
	if y >= height || x >= width {
		return 0
	}
	return uint32(y*width + x)
}
