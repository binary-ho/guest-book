package common

import (
	"database/sql/driver"
	"errors"
	"net/url"
)

type Url struct {
	url *url.URL
}

func CreateUrl(rawURl string) (Url, error) {
	parsedUrl, err := url.Parse(rawURl)
	if err != nil {
		return Url{}, err
	}
	return Url{parsedUrl}, err
}

func (u Url) Value() (driver.Value, error) {
	if u.url == nil {
		return nil, errors.New("Url is nil!")
	}
	return u.url.String(), nil
}

func (u *Url) Scan(value interface{}) error {
	result, ok := value.([]byte)
	if !ok {
		return errors.New("url을 []byte type으로 형변환 불가능")
	}

	str := string(result[:])
	parsedUrl, err := url.Parse(str)
	if err != nil {
		return err
	}
	u.url = parsedUrl
	return nil
}
