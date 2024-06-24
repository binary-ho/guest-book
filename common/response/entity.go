package response

import (
	"encoding/json"
	"errors"
	"guestbook/common/util"
	"net/http"
	"strconv"
)

const (
	ContentTypeKey  = "Content-Type"
	JsonContentType = "application/json"
)

func Create(writer *http.ResponseWriter, responseBody any, status int) error {
	err := json.NewEncoder(*writer).Encode(responseBody)
	if err != nil {
		Error(*writer, http.StatusNotFound)
		return err
	}

	(*writer).Header().Set(ContentTypeKey, JsonContentType)
	(*writer).WriteHeader(status)
	err = validateStatusNumber(status)
	return err
}

func validateStatusNumber(status int) error {
	text := http.StatusText(status)
	if util.IsBlank(text) {
		return errors.New("Wrong Status Code : " + strconv.Itoa(status))
	}
	return nil
}
