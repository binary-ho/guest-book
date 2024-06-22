package main

import (
	"guestbook/common/util"
	"net/http"
	"strconv"
)

func (app *application) internal(writer http.ResponseWriter, request *http.Request, err error) {
	var (
		method = request.Method
		uri    = request.URL.RequestURI()
	)

	app.logger.Error("[INTERNAL ERROR]",
		"error", err.Error(), "method", method, "uri", uri)
	internalError := http.StatusInternalServerError
	http.Error(writer, http.StatusText(internalError), internalError)
}

func (app *application) client(writer http.ResponseWriter, status int) {
	statusText := http.StatusText(status)
	if util.IsBlank(statusText) {
		status = http.StatusNotFound
		app.logger.Error("[CLIENT ERROR - STATUS WRONG]", "status", strconv.Itoa(status))
	}
	app.logger.Error("[CLIENT ERROR] status", "status", strconv.Itoa(status))
	http.Error(writer, statusText, status)
}