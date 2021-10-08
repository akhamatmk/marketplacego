package application

import (
	"log"
	"marketplacego/config"
	"marketplacego/server"
	"marketplacego/server/routes"
)

func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	routes.ConfigureRoutes(app)

	// err := app.Start(cfg.HTTP.Port)
	err := app.Start("5000")
	if err != nil {
		log.Fatal("Port already used")
	}
}
