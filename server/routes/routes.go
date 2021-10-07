package routes

import (
	s "marketplacego/server"
	"marketplacego/services/token"

	"github.com/labstack/echo/v4/middleware"
	// "echo-demo-project/server/handlers"
	// "echo-demo-project/services/token"
)

func ConfigureRoutes(server *s.Server) {
	// postHandler := handlers.NewPostHandlers(server)
	// authHandler := handlers.NewAuthHandler(server)
	// registerHandler := handlers.NewRegisterHandler(server)

	server.Echo.Use(middleware.Logger())
	// server.Echo.POST("/login", authHandler.Login)
	// server.Echo.POST("/register", registerHandler.Register)
	// server.Echo.POST("/refresh", authHandler.RefreshToken)

	// fmt.Println(server.Config.Auth.AccessSecret)

	r := server.Echo.Group("")
	config := middleware.JWTConfig{
		Claims:     &token.JwtCustomClaims{},
		SigningKey: []byte(server.Config.Auth.AccessSecret),
	}
	r.Use(middleware.JWTWithConfig(config))

	// r.GET("/posts", postHandler.GetPosts)
	// r.POST("/posts", postHandler.CreatePost)
	// r.DELETE("/posts/:id", postHandler.DeletePost)
	// r.PUT("/posts/:id", postHandler.UpdatePost)
}
