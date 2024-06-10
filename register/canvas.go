package register

import "guestbook/entity"

type Canvas struct {
	id   entity.ID
	data map[string]LWWRegister
}

func (llwMap Canvas) IsExist(key string) bool {
	_, exist := llwMap.data[key]
	return exist
}
