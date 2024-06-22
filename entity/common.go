package entity

import (
	"database/sql/driver"
	"errors"
	"strconv"
)

type ID uint

func DefaultId() ID {
	return ID(0)
}

func (id ID) Value() (driver.Value, error) {
	return int64(id), nil
}

func (id *ID) Scan(value interface{}) error {

	switch result := value.(type) {
	case nil:
		return nil
	case int64:
		*id = ID(result)
	case []byte:
		number, err := strconv.ParseUint(string(result), 10, 64)
		if err != nil {
			return err
		}
		*id = ID(number)
	default:
		return errors.New("ID로 형변환 불가능 ")
	}
	return nil
}
