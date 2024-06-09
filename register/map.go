package register

import "guestbook/entity"

type RGBMap struct {
	id   entity.Id
	data map[string]LWWRegister
}

func (llwMap RGBMap) Has(key string) bool {
	_, exist := llwMap.data[key]
	return exist
}
