package entity

import (
	"errors"
)

type ID uint

func DefaultId() ID {
	return ID(0)
}

func (id ID) Value() uint {
	return uint(id)
}

func (id *ID) Scan(value interface{}) error {
	result, ok := value.(uint)
	if !ok {
		return errors.New("uint type으로 형변환 불가능")
	}
	*id = ID(result)
	return nil
}
