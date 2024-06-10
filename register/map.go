package register

import "guestbook/entity"

type RGBMap struct {
	id   entity.ID
	data map[string]LWWRegister
}

func (llwMap RGBMap) Has(key string) bool {
	_, exist := llwMap.data[key]
	return exist
}
