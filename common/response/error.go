package response

import (
	"net/http"
)

func Error(writer http.ResponseWriter, status int) {
	http.Error(writer, http.StatusText(status), status)
}
