package routes

import (
	s "marketplacego/server"
	"marketplacego/server/handlers"
	"marketplacego/services/token"

	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(server *s.Server) {
	registerHandler := handlers.NewRegisterHandler(server)

	server.Echo.Use(middleware.Logger())
	server.Echo.POST("/user/register", registerHandler.Register)

	// fmt.Println(server.Config.Auth.AccessSecret)

	r := server.Echo.Group("")
	config := middleware.JWTConfig{
		Claims:     &token.JwtCustomClaims{},
		SigningKey: []byte(server.Config.Auth.AccessSecret),
	}
	r.Use(middleware.JWTWithConfig(config))
}
