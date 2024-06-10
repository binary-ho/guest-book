package register

import "guestbook/entity"

type Canvas struct {
	id   entity.ID
	data map[uint32]LWWRegister
}

func (llwMap Canvas) IsExist(key uint32) bool {
	_, exist := llwMap.data[key]
	return exist
}
