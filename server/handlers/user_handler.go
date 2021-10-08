package handlers

import (
	"marketplacego/models"
	"marketplacego/repositories"
	"marketplacego/responses"
	s "marketplacego/server"
	tokenservice "marketplacego/services/token"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	server *s.Server
}

func NewUserHandler(server *s.Server) *UserHandler {
	return &UserHandler{server: server}
}

func (userHandler *UserHandler) TestLogin(c echo.Context) error {

	user := models.User{}
	userRepository := repositories.NewUserRepository(userHandler.server.DB)
	userRepository.GetUserByEmail(&user, "akham@gmail.com")

	tokenService := tokenservice.NewTokenService(userHandler.server.Config)
	accessToken, exp, err := tokenService.CreateAccessToken(&user)
	if err != nil {
		return err
	}
	refreshToken, err := tokenService.CreateRefreshToken(&user)
	if err != nil {
		return err
	}
	res := responses.NewLoginResponse(accessToken, refreshToken, exp)

	return responses.MessageResponse(c, http.StatusOK, "Succes login", res)
}
