package user

import (
	"encoding/json"
	"guestbook/common/response"
	"guestbook/entity"
	user "guestbook/entity/user/dto"
	"net/http"
)

type Controller struct {
	service *Service
}

func InitController(service *Service) *Controller {
	return &Controller{service: service}
}

func (controller *Controller) SignUp() (string, func(http.ResponseWriter, *http.Request)) {
	SignUpEndpoint := "POST /user"
	return SignUpEndpoint, controller.signUp
}

func (controller *Controller) signUp(writer http.ResponseWriter, request *http.Request) {
	var signUpRequest user.SignUpRequest
	err := json.NewDecoder(request.Body).Decode(&signUpRequest)
	if err != nil {
		response.Error(writer, http.StatusNotFound)
		return
	}
	service := controller.service
	githubId := entity.ID(signUpRequest.GithubId)

	id, err := service.SignUp(signUpRequest.Handle, githubId, signUpRequest.ProfileUrl)
	if err != nil {
		response.Error(writer, http.StatusNotFound)
		return
	}

	err = response.Create(&writer, id, http.StatusCreated)
	if err != nil {
		response.Error(writer, http.StatusNotFound)
		return
	}
}
