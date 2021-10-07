package main

import (
	application "marketplacego"
	"marketplacego/config"
)

func main() {
	cfg := config.NewConfig()
	application.Start(cfg)
}
