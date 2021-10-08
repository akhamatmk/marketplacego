package handlers

import (
	"marketplacego/models"
	"marketplacego/repositories"
	"marketplacego/requests"
	"marketplacego/responses"
	s "marketplacego/server"
	"marketplacego/services/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RegisterHandler struct {
	server *s.Server
}

func NewRegisterHandler(server *s.Server) *RegisterHandler {
	return &RegisterHandler{server: server}
}

func (registerHandler *RegisterHandler) Register(c echo.Context) error {
	registerRequest := new(requests.RegisterRequest)

	if err := c.Bind(registerRequest); err != nil {
		return err
	}

	if err := registerRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid")
	}

	existUser := models.User{}
	userRepository := repositories.NewUserRepository(registerHandler.server.DB)
	userRepository.GetUserByEmail(&existUser, registerRequest.Email)

	if existUser.ID != 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "User already exists")
	}

	userService := user.NewUserService(registerHandler.server.DB)
	if err := userService.Register(registerRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Server error")
	}

	return responses.MessageResponse(c, http.StatusCreated, "User successfully created", nil)
}
