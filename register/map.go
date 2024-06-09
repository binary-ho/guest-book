package register

type RGBMap struct {
	id   string
	data map[string]LWWRegister
}

func (llwMap RGBMap) Has(key string) bool {
	_, exist := llwMap.data[key]
	return exist
}
